package common_test

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"

	"code.cloudfoundry.org/cli/actor/pluginaction"
	"code.cloudfoundry.org/cli/api/plugin/pluginerror"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/commandfakes"
	. "code.cloudfoundry.org/cli/command/common"
	"code.cloudfoundry.org/cli/command/common/commonfakes"
	"code.cloudfoundry.org/cli/command/plugin/shared"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("install-plugin command", func() {
	var (
		cmd         InstallPluginCommand
		testUI      *ui.UI
		input       *Buffer
		fakeConfig  *commandfakes.FakeConfig
		fakeActor   *commonfakes.FakeInstallPluginActor
		executeErr  error
		expectedErr error
		pluginHome  string
	)

	BeforeEach(func() {
		input = NewBuffer()
		testUI = ui.NewTestUI(input, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeActor = new(commonfakes.FakeInstallPluginActor)

		cmd = InstallPluginCommand{
			UI:     testUI,
			Config: fakeConfig,
			Actor:  fakeActor,
		}

		tmpDirectorySeed := strconv.Itoa(int(rand.Int63()))
		pluginHome = fmt.Sprintf("some-pluginhome-%s", tmpDirectorySeed)
		fakeConfig.PluginHomeReturns(pluginHome)
		fakeConfig.ExperimentalReturns(true)
		fakeConfig.BinaryNameReturns("faceman")
	})

	AfterEach(func() {
		os.RemoveAll(pluginHome)
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	Describe("installing from a local file", func() {
		BeforeEach(func() {
			cmd.OptionalArgs.PluginNameOrLocation = "some-path"
		})

		Context("when the local file does not exist", func() {
			BeforeEach(func() {
				fakeActor.FileExistsReturns(false)
			})

			It("does not print installation messages and returns a FileNotFoundError", func() {
				Expect(executeErr).To(MatchError(shared.FileNotFoundError{Path: "some-path"}))

				Expect(testUI.Out).ToNot(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
				Expect(testUI.Out).ToNot(Say("Installing plugin some-path\\.\\.\\."))
			})
		})

		Context("when the file exists", func() {
			BeforeEach(func() {
				fakeActor.CreateExecutableCopyReturns("copy-path", nil)
				fakeActor.FileExistsReturns(true)
			})

			Context("when the -f argument is given", func() {
				BeforeEach(func() {
					cmd.Force = true
				})

				Context("when the plugin is invalid", func() {
					var returnedErr error

					BeforeEach(func() {
						returnedErr = pluginaction.PluginInvalidError{}
						fakeActor.GetAndValidatePluginReturns(configv3.Plugin{}, returnedErr)
					})

					It("returns an error", func() {
						Expect(executeErr).To(MatchError(shared.PluginInvalidError{}))

						Expect(testUI.Out).ToNot(Say("Installing plugin"))
					})
				})

				Context("when the plugin is already installed", func() {
					var plugin configv3.Plugin

					BeforeEach(func() {
						plugin = configv3.Plugin{
							Name: "some-plugin",
							Version: configv3.PluginVersion{
								Major: 1,
								Minor: 2,
								Build: 3,
							},
						}
						fakeActor.GetAndValidatePluginReturns(plugin, nil)
						fakeActor.IsPluginInstalledReturns(true)
					})

					Context("when an error is encountered uninstalling the existing plugin", func() {
						BeforeEach(func() {
							expectedErr = errors.New("uninstall plugin error")
							fakeActor.UninstallPluginReturns(expectedErr)
						})

						It("returns the error", func() {
							Expect(executeErr).To(MatchError(expectedErr))

							Expect(testUI.Out).ToNot(Say("Plugin some-plugin successfully uninstalled\\."))
						})
					})

					Context("when no errors are encountered uninstalling the existing plugin", func() {
						It("uninstalls the existing plugin and installs the current plugin", func() {
							Expect(executeErr).ToNot(HaveOccurred())

							Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
							Expect(testUI.Out).To(Say("Install and use plugins at your own risk\\."))
							Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.3 is already installed\\. Uninstalling existing plugin\\.\\.\\."))
							Expect(testUI.Out).To(Say("OK"))
							Expect(testUI.Out).To(Say("Plugin some-plugin successfully uninstalled\\."))
							Expect(testUI.Out).To(Say("Installing plugin some-plugin\\.\\.\\."))
							Expect(testUI.Out).To(Say("OK"))
							Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.3 successfully installed\\."))

							Expect(fakeActor.FileExistsCallCount()).To(Equal(1))
							Expect(fakeActor.FileExistsArgsForCall(0)).To(Equal("some-path"))

							Expect(fakeActor.GetAndValidatePluginCallCount()).To(Equal(1))
							_, _, path := fakeActor.GetAndValidatePluginArgsForCall(0)
							Expect(path).To(Equal("copy-path"))

							Expect(fakeActor.IsPluginInstalledCallCount()).To(Equal(1))
							Expect(fakeActor.IsPluginInstalledArgsForCall(0)).To(Equal("some-plugin"))

							Expect(fakeActor.UninstallPluginCallCount()).To(Equal(1))
							_, pluginName := fakeActor.UninstallPluginArgsForCall(0)
							Expect(pluginName).To(Equal("some-plugin"))

							Expect(fakeActor.InstallPluginFromPathCallCount()).To(Equal(1))
							path, installedPlugin := fakeActor.InstallPluginFromPathArgsForCall(0)
							Expect(path).To(Equal("copy-path"))
							Expect(installedPlugin).To(Equal(plugin))
						})

						Context("when an error is encountered installing the plugin", func() {
							BeforeEach(func() {
								expectedErr = errors.New("install plugin error")
								fakeActor.InstallPluginFromPathReturns(expectedErr)
							})

							It("returns the error", func() {
								Expect(executeErr).To(MatchError(expectedErr))

								Expect(testUI.Out).ToNot(Say("Plugin some-plugin 1\\.2\\.3 successfully installed\\."))
							})
						})
					})
				})

				Context("when the plugin is not already installed", func() {
					var plugin configv3.Plugin

					BeforeEach(func() {
						plugin = configv3.Plugin{
							Name: "some-plugin",
							Version: configv3.PluginVersion{
								Major: 1,
								Minor: 2,
								Build: 3,
							},
						}
						fakeActor.GetAndValidatePluginReturns(plugin, nil)
					})

					It("installs the plugin", func() {
						Expect(executeErr).ToNot(HaveOccurred())

						Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Expect(testUI.Out).To(Say("Install and use plugins at your own risk\\."))
						Expect(testUI.Out).To(Say("Installing plugin some-plugin\\.\\.\\."))
						Expect(testUI.Out).To(Say("OK"))
						Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.3 successfully installed\\."))

						Expect(fakeActor.FileExistsCallCount()).To(Equal(1))
						Expect(fakeActor.FileExistsArgsForCall(0)).To(Equal("some-path"))

						Expect(fakeActor.CreateExecutableCopyCallCount()).To(Equal(1))
						pathArg, pluginDirArg := fakeActor.CreateExecutableCopyArgsForCall(0)
						Expect(pathArg).To(Equal("some-path"))
						Expect(pluginDirArg).To(ContainSubstring("some-pluginhome"))
						Expect(pluginDirArg).To(ContainSubstring("temp"))

						Expect(fakeActor.GetAndValidatePluginCallCount()).To(Equal(1))
						_, _, path := fakeActor.GetAndValidatePluginArgsForCall(0)
						Expect(path).To(Equal("copy-path"))

						Expect(fakeActor.IsPluginInstalledCallCount()).To(Equal(1))
						Expect(fakeActor.IsPluginInstalledArgsForCall(0)).To(Equal("some-plugin"))

						Expect(fakeActor.InstallPluginFromPathCallCount()).To(Equal(1))
						path, installedPlugin := fakeActor.InstallPluginFromPathArgsForCall(0)
						Expect(path).To(Equal("copy-path"))
						Expect(installedPlugin).To(Equal(plugin))

						Expect(fakeActor.UninstallPluginCallCount()).To(Equal(0))
					})

					Context("when there is an error making an executable copy of the plugin binary", func() {
						BeforeEach(func() {
							expectedErr = errors.New("create executable copy error")
							fakeActor.CreateExecutableCopyReturns("", expectedErr)
						})

						It("returns the error", func() {
							Expect(executeErr).To(MatchError(expectedErr))
						})
					})

					Context("when an error is encountered installing the plugin", func() {
						BeforeEach(func() {
							expectedErr = errors.New("install plugin error")
							fakeActor.InstallPluginFromPathReturns(expectedErr)
						})

						It("returns the error", func() {
							Expect(executeErr).To(MatchError(expectedErr))

							Expect(testUI.Out).ToNot(Say("Plugin some-plugin 1\\.2\\.3 successfully installed\\."))
						})
					})
				})
			})

			Context("when the -f argument is not given (user is prompted for confirmation)", func() {
				BeforeEach(func() {
					cmd.Force = false
				})

				Context("when the user chooses no", func() {
					BeforeEach(func() {
						input.Write([]byte("n\n"))
					})

					It("cancels plugin installation", func() {
						Expect(executeErr).ToNot(HaveOccurred())

						Expect(testUI.Out).To(Say("Plugin installation cancelled\\."))
					})
				})

				Context("when the user chooses the default", func() {
					BeforeEach(func() {
						input.Write([]byte("\n"))
					})

					It("cancels plugin installation", func() {
						Expect(executeErr).ToNot(HaveOccurred())

						Expect(testUI.Out).To(Say("Plugin installation cancelled\\."))
					})
				})

				Context("when the user input is invalid", func() {
					BeforeEach(func() {
						input.Write([]byte("e\n"))
					})

					It("returns an error", func() {
						Expect(executeErr).To(HaveOccurred())

						Expect(testUI.Out).ToNot(Say("Installing plugin"))
					})
				})

				Context("when the user chooses yes", func() {
					BeforeEach(func() {
						input.Write([]byte("y\n"))
					})

					Context("when the plugin is not already installed", func() {
						var plugin configv3.Plugin

						BeforeEach(func() {
							plugin = configv3.Plugin{
								Name: "some-plugin",
								Version: configv3.PluginVersion{
									Major: 1,
									Minor: 2,
									Build: 3,
								},
							}
							fakeActor.GetAndValidatePluginReturns(plugin, nil)
						})

						It("installs the plugin", func() {
							Expect(executeErr).ToNot(HaveOccurred())

							Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
							Expect(testUI.Out).To(Say("Install and use plugins at your own risk\\."))
							Expect(testUI.Out).To(Say("Do you want to install the plugin some-path\\? \\[yN\\]"))
							Expect(testUI.Out).To(Say("Installing plugin some-plugin\\.\\.\\."))
							Expect(testUI.Out).To(Say("OK"))
							Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.3 successfully installed\\."))

							Expect(fakeActor.FileExistsCallCount()).To(Equal(1))
							Expect(fakeActor.FileExistsArgsForCall(0)).To(Equal("some-path"))

							Expect(fakeActor.GetAndValidatePluginCallCount()).To(Equal(1))
							_, _, path := fakeActor.GetAndValidatePluginArgsForCall(0)
							Expect(path).To(Equal("copy-path"))

							Expect(fakeActor.IsPluginInstalledCallCount()).To(Equal(1))
							Expect(fakeActor.IsPluginInstalledArgsForCall(0)).To(Equal("some-plugin"))

							Expect(fakeActor.InstallPluginFromPathCallCount()).To(Equal(1))
							path, plugin := fakeActor.InstallPluginFromPathArgsForCall(0)
							Expect(path).To(Equal("copy-path"))
							Expect(plugin).To(Equal(plugin))

							Expect(fakeActor.UninstallPluginCallCount()).To(Equal(0))
						})
					})

					Context("when the plugin is already installed", func() {
						BeforeEach(func() {
							plugin := configv3.Plugin{
								Name: "some-plugin",
								Version: configv3.PluginVersion{
									Major: 1,
									Minor: 2,
									Build: 3,
								},
							}
							fakeActor.GetAndValidatePluginReturns(plugin, nil)
							fakeActor.IsPluginInstalledReturns(true)
						})

						It("returns PluginAlreadyInstalledError", func() {
							Expect(executeErr).To(MatchError(shared.PluginAlreadyInstalledError{
								BinaryName: "faceman",
								Name:       "some-plugin",
								Version:    "1.2.3",
							}))
						})
					})
				})
			})
		})
	})

	Describe("installing from an unsupported URL scheme", func() {
		BeforeEach(func() {
			cmd.OptionalArgs.PluginNameOrLocation = "ftp://some-url"
		})

		It("returns an error indicating an unsupported URL scheme", func() {
			Expect(executeErr).To(MatchError(command.UnsupportedURLSchemeError{
				UnsupportedURL: string(cmd.OptionalArgs.PluginNameOrLocation),
			}))
		})
	})

	Describe("installing from an HTTP URL", func() {
		var (
			plugin               configv3.Plugin
			pluginName           string
			downloadedPluginPath string
			executablePluginPath string
		)

		BeforeEach(func() {
			cmd.OptionalArgs.PluginNameOrLocation = "http://some-url"
			pluginName = "some-plugin"
			downloadedPluginPath = "some-path"
			executablePluginPath = "executable-path"
		})

		It("displays the plugin warning", func() {
			Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
			Expect(testUI.Out).To(Say("Install and use plugins at your own risk\\."))
		})

		Context("when the -f argument is given", func() {
			BeforeEach(func() {
				cmd.Force = true
			})

			It("begins downloading the plugin", func() {
				Expect(testUI.Out).To(Say("Starting download of plugin binary from URL\\.\\.\\."))

				Expect(fakeActor.DownloadExecutableBinaryFromURLCallCount()).To(Equal(1))
				url, tempPluginDir := fakeActor.DownloadExecutableBinaryFromURLArgsForCall(0)
				Expect(url).To(Equal(cmd.OptionalArgs.PluginNameOrLocation.String()))
				Expect(tempPluginDir).To(ContainSubstring("some-pluginhome"))
				Expect(tempPluginDir).To(ContainSubstring("temp"))
			})

			Context("When getting the binary fails", func() {
				BeforeEach(func() {
					expectedErr = errors.New("some-error")
					fakeActor.DownloadExecutableBinaryFromURLReturns("", 0, expectedErr)
				})

				It("returns the error", func() {
					Expect(executeErr).To(MatchError(expectedErr))

					Expect(testUI.Out).ToNot(Say("downloaded"))
					Expect(fakeActor.GetAndValidatePluginCallCount()).To(Equal(0))
				})

				Context("when a 4xx or 5xx status is encountered while downloading the plugin", func() {
					BeforeEach(func() {
						fakeActor.DownloadExecutableBinaryFromURLReturns("", 0, pluginerror.RawHTTPStatusError{Status: "some-status"})
					})

					It("returns a DownloadPluginHTTPError", func() {
						Expect(executeErr).To(MatchError(shared.DownloadPluginHTTPError{Message: "some-status"}))
					})
				})

				Context("when a SSL error is encountered while downloading the plugin", func() {
					BeforeEach(func() {
						fakeActor.DownloadExecutableBinaryFromURLReturns("", 0, pluginerror.UnverifiedServerError{})
					})

					It("returns a DownloadPluginHTTPError", func() {
						Expect(executeErr).To(MatchError(shared.DownloadPluginHTTPError{Message: "x509: certificate signed by unknown authority"}))
					})
				})
			})

			Context("when getting the binary succeeds", func() {
				BeforeEach(func() {
					fakeActor.DownloadExecutableBinaryFromURLReturns("some-path", 4, nil)
					fakeActor.CreateExecutableCopyReturns(executablePluginPath, nil)
				})

				It("displays the bytes downloaded", func() {
					Expect(testUI.Out).To(Say("4 bytes downloaded\\.\\.\\."))

					Expect(fakeActor.GetAndValidatePluginCallCount()).To(Equal(1))
					_, _, path := fakeActor.GetAndValidatePluginArgsForCall(0)
					Expect(path).To(Equal(executablePluginPath))

					Expect(fakeActor.DownloadExecutableBinaryFromURLCallCount()).To(Equal(1))
					urlArg, pluginDirArg := fakeActor.DownloadExecutableBinaryFromURLArgsForCall(0)
					Expect(urlArg).To(Equal("http://some-url"))
					Expect(pluginDirArg).To(ContainSubstring("some-pluginhome"))
					Expect(pluginDirArg).To(ContainSubstring("temp"))

					Expect(fakeActor.CreateExecutableCopyCallCount()).To(Equal(1))
					pathArg, pluginDirArg := fakeActor.CreateExecutableCopyArgsForCall(0)
					Expect(pathArg).To(Equal("some-path"))
					Expect(pluginDirArg).To(ContainSubstring("some-pluginhome"))
					Expect(pluginDirArg).To(ContainSubstring("temp"))
				})

				Context("when the plugin is invalid", func() {
					var returnedErr error

					BeforeEach(func() {
						returnedErr = pluginaction.PluginInvalidError{}
						fakeActor.GetAndValidatePluginReturns(configv3.Plugin{}, returnedErr)
					})

					It("returns an error", func() {
						Expect(executeErr).To(MatchError(shared.PluginInvalidError{}))

						Expect(fakeActor.IsPluginInstalledCallCount()).To(Equal(0))
					})
				})

				Context("when the plugin is valid", func() {
					BeforeEach(func() {
						plugin = configv3.Plugin{
							Name: pluginName,
							Version: configv3.PluginVersion{
								Major: 1,
								Minor: 2,
								Build: 3,
							},
						}
						fakeActor.GetAndValidatePluginReturns(plugin, nil)
					})

					Context("when the plugin is already installed", func() {
						BeforeEach(func() {
							fakeActor.IsPluginInstalledReturns(true)
						})

						It("displays uninstall message", func() {
							Expect(testUI.Out).To(Say("Plugin %s 1\\.2\\.3 is already installed\\. Uninstalling existing plugin\\.\\.\\.", pluginName))
						})

						Context("when an error is encountered uninstalling the existing plugin", func() {
							BeforeEach(func() {
								expectedErr = errors.New("uninstall plugin error")
								fakeActor.UninstallPluginReturns(expectedErr)
							})

							It("returns the error", func() {
								Expect(executeErr).To(MatchError(expectedErr))

								Expect(testUI.Out).ToNot(Say("Plugin some-plugin successfully uninstalled\\."))
							})
						})

						Context("when no errors are encountered uninstalling the existing plugin", func() {
							It("displays uninstall message", func() {
								Expect(testUI.Out).To(Say("Plugin %s successfully uninstalled\\.", pluginName))
							})

							Context("when no errors are encountered installing the plugin", func() {
								It("uninstalls the existing plugin and installs the current plugin", func() {
									Expect(executeErr).ToNot(HaveOccurred())

									Expect(testUI.Out).To(Say("Installing plugin %s\\.\\.\\.", pluginName))
									Expect(testUI.Out).To(Say("OK"))
									Expect(testUI.Out).To(Say("Plugin %s 1\\.2\\.3 successfully installed\\.", pluginName))
								})
							})

							Context("when an error is encountered installing the plugin", func() {
								BeforeEach(func() {
									expectedErr = errors.New("install plugin error")
									fakeActor.InstallPluginFromPathReturns(expectedErr)
								})

								It("returns the error", func() {
									Expect(executeErr).To(MatchError(expectedErr))

									Expect(testUI.Out).ToNot(Say("Plugin some-plugin 1\\.2\\.3 successfully installed\\."))
								})
							})
						})
					})

					Context("when the plugin is not already installed", func() {
						It("installs the plugin", func() {
							Expect(executeErr).ToNot(HaveOccurred())

							Expect(testUI.Out).To(Say("Installing plugin %s\\.\\.\\.", pluginName))
							Expect(testUI.Out).To(Say("OK"))
							Expect(testUI.Out).To(Say("Plugin %s 1\\.2\\.3 successfully installed\\.", pluginName))

							Expect(fakeActor.UninstallPluginCallCount()).To(Equal(0))
						})
					})
				})
			})
		})

		Context("when the -f argument is not given (user is prompted for confirmation)", func() {
			BeforeEach(func() {
				plugin = configv3.Plugin{
					Name: pluginName,
					Version: configv3.PluginVersion{
						Major: 1,
						Minor: 2,
						Build: 3,
					},
				}

				cmd.Force = false
				fakeActor.DownloadExecutableBinaryFromURLReturns("some-path", 4, nil)
				fakeActor.CreateExecutableCopyReturns("executable-path", nil)
			})

			Context("when the user chooses no", func() {
				BeforeEach(func() {
					input.Write([]byte("n\n"))
				})

				It("cancels plugin installation", func() {
					Expect(executeErr).ToNot(HaveOccurred())

					Expect(testUI.Out).To(Say("Plugin installation cancelled\\."))
				})
			})

			Context("when the user chooses the default", func() {
				BeforeEach(func() {
					input.Write([]byte("\n"))
				})

				It("cancels plugin installation", func() {
					Expect(executeErr).ToNot(HaveOccurred())

					Expect(testUI.Out).To(Say("Plugin installation cancelled\\."))
				})
			})

			Context("when the user input is invalid", func() {
				BeforeEach(func() {
					input.Write([]byte("e\n"))
				})

				It("returns an error", func() {
					Expect(executeErr).To(HaveOccurred())

					Expect(testUI.Out).ToNot(Say("Installing plugin"))
				})
			})

			Context("when the user chooses yes", func() {
				BeforeEach(func() {
					input.Write([]byte("y\n"))
				})

				Context("when the plugin is not already installed", func() {
					BeforeEach(func() {
						fakeActor.GetAndValidatePluginReturns(plugin, nil)
					})

					It("installs the plugin", func() {
						Expect(executeErr).ToNot(HaveOccurred())

						Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Expect(testUI.Out).To(Say("Install and use plugins at your own risk\\."))
						Expect(testUI.Out).To(Say("Do you want to install the plugin %s\\? \\[yN\\]", cmd.OptionalArgs.PluginNameOrLocation))
						Expect(testUI.Out).To(Say("Starting download of plugin binary from URL\\.\\.\\."))

						Expect(testUI.Out).To(Say("4 bytes downloaded\\.\\.\\."))
						Expect(testUI.Out).To(Say("Installing plugin %s\\.\\.\\.", pluginName))
						Expect(testUI.Out).To(Say("OK"))
						Expect(testUI.Out).To(Say("Plugin %s 1\\.2\\.3 successfully installed\\.", pluginName))

						Expect(fakeActor.DownloadExecutableBinaryFromURLCallCount()).To(Equal(1))
						url, tempPluginDir := fakeActor.DownloadExecutableBinaryFromURLArgsForCall(0)
						Expect(url).To(Equal(cmd.OptionalArgs.PluginNameOrLocation.String()))
						Expect(tempPluginDir).To(ContainSubstring("some-pluginhome"))
						Expect(tempPluginDir).To(ContainSubstring("temp"))

						Expect(fakeActor.CreateExecutableCopyCallCount()).To(Equal(1))
						path, tempPluginDir := fakeActor.CreateExecutableCopyArgsForCall(0)
						Expect(path).To(Equal("some-path"))
						Expect(tempPluginDir).To(ContainSubstring("some-pluginhome"))
						Expect(tempPluginDir).To(ContainSubstring("temp"))

						Expect(fakeActor.GetAndValidatePluginCallCount()).To(Equal(1))
						_, _, path = fakeActor.GetAndValidatePluginArgsForCall(0)
						Expect(path).To(Equal(executablePluginPath))

						Expect(fakeActor.IsPluginInstalledCallCount()).To(Equal(1))
						Expect(fakeActor.IsPluginInstalledArgsForCall(0)).To(Equal(pluginName))

						Expect(fakeActor.InstallPluginFromPathCallCount()).To(Equal(1))
						path, installedPlugin := fakeActor.InstallPluginFromPathArgsForCall(0)
						Expect(path).To(Equal(executablePluginPath))
						Expect(installedPlugin).To(Equal(plugin))

						Expect(fakeActor.UninstallPluginCallCount()).To(Equal(0))
					})
				})

				Context("when the plugin is already installed", func() {
					BeforeEach(func() {
						fakeActor.GetAndValidatePluginReturns(plugin, nil)
						fakeActor.IsPluginInstalledReturns(true)
					})

					It("returns PluginAlreadyInstalledError", func() {
						Expect(executeErr).To(MatchError(shared.PluginAlreadyInstalledError{
							BinaryName: "faceman",
							Name:       pluginName,
							Version:    "1.2.3",
						}))
					})
				})
			})
		})
	})

	Describe("installing from a specific repo", func() {
		BeforeEach(func() {
			cmd.OptionalArgs.PluginNameOrLocation = "some-plugin"
			cmd.RegisteredRepository = "some-repo"
		})

		Context("when the repo is not registered", func() {
			BeforeEach(func() {
				fakeActor.GetPluginRepositoryReturns(configv3.PluginRepository{}, pluginaction.RepositoryNotRegisteredError{Name: "some-repo"})
			})

			It("returns a RepositoryNotRegisteredError", func() {
				Expect(executeErr).To(MatchError(shared.RepositoryNotRegisteredError{Name: "some-repo"}))

				Expect(fakeActor.GetPluginRepositoryCallCount()).To(Equal(1))
				repositoryNameArg := fakeActor.GetPluginRepositoryArgsForCall(0)
				Expect(repositoryNameArg).To(Equal("some-repo"))
			})
		})

		Context("when the repository is registered", func() {
			BeforeEach(func() {
				fakeActor.GetPluginRepositoryReturns(configv3.PluginRepository{Name: "some-repo", URL: "http://some-url"}, nil)
			})

			Context("when the plugin can't be found in the repository", func() {
				BeforeEach(func() {
					fakeActor.GetPlatformStringReturns("platform-i-dont-exist")
					fakeActor.GetPluginInfoFromRepositoryForPlatformReturns(pluginaction.PluginInfo{}, pluginaction.PluginNotFoundInRepositoryError{PluginName: "some-plugin", RepositoryName: "some-repo"})
				})

				It("returns the PluginNotFoundInRepositoryError", func() {
					Expect(executeErr).To(MatchError(shared.PluginNotFoundInRepositoryError{BinaryName: "faceman", PluginName: "some-plugin", RepositoryName: "some-repo"}))

					Expect(fakeActor.GetPlatformStringCallCount()).To(Equal(1))
					platformGOOS, platformGOARCH := fakeActor.GetPlatformStringArgsForCall(0)
					Expect(platformGOOS).To(Equal(runtime.GOOS))
					Expect(platformGOARCH).To(Equal(runtime.GOARCH))

					Expect(fakeActor.GetPluginInfoFromRepositoryForPlatformCallCount()).To(Equal(1))
					pluginNameArg, pluginRepositoryArg, pluginPlatform := fakeActor.GetPluginInfoFromRepositoryForPlatformArgsForCall(0)
					Expect(pluginNameArg).To(Equal("some-plugin"))
					Expect(pluginRepositoryArg).To(Equal(configv3.PluginRepository{Name: "some-repo", URL: "http://some-url"}))
					Expect(pluginPlatform).To(Equal("platform-i-dont-exist"))
				})
			})

			Context("when the plugin is found", func() {
				BeforeEach(func() {
					fakeActor.GetPlatformStringReturns("linux64")
					fakeActor.GetPluginInfoFromRepositoryForPlatformReturns(pluginaction.PluginInfo{
						Name:     "some-plugin",
						Version:  "1.2.3",
						URL:      "http://some-url",
						Checksum: "some-checksum",
					}, nil)
				})

				Context("when the -f argument is given", func() {
					BeforeEach(func() {
						cmd.Force = true
					})

					Context("when the plugin is already installed", func() {
						BeforeEach(func() {
							fakeConfig.GetPluginReturns(configv3.Plugin{
								Name:    "some-plugin",
								Version: configv3.PluginVersion{Major: 1, Minor: 2, Build: 2},
							}, true)
							fakeActor.IsPluginInstalledReturns(true)
						})

						Context("when getting the binary errors", func() {
							BeforeEach(func() {
								expectedErr = errors.New("some-error")
								fakeActor.DownloadExecutableBinaryFromURLReturns("", 0, expectedErr)
							})

							It("returns the error", func() {
								Expect(executeErr).To(MatchError(expectedErr))

								Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin..."))
								Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.3 found in: some-repo"))
								Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.2 is already installed."))
								Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors."))
								Expect(testUI.Out).To(Say("Install and use plugins at your own risk."))
								Expect(testUI.Out).To(Say("Starting download of plugin binary from repository some-repo..."))

								Expect(testUI.Out).ToNot(Say("downloaded"))
								Expect(fakeActor.GetAndValidatePluginCallCount()).To(Equal(0))

								Expect(fakeConfig.GetPluginCallCount()).To(Equal(1))
								Expect(fakeConfig.GetPluginArgsForCall(0)).To(Equal("some-plugin"))

								Expect(fakeActor.GetPlatformStringCallCount()).To(Equal(1))
								platformGOOS, platformGOARCH := fakeActor.GetPlatformStringArgsForCall(0)
								Expect(platformGOOS).To(Equal(runtime.GOOS))
								Expect(platformGOARCH).To(Equal(runtime.GOARCH))

								Expect(fakeActor.GetPluginInfoFromRepositoryForPlatformCallCount()).To(Equal(1))
								pluginNameArg, pluginRepositoryArg, pluginPlatform := fakeActor.GetPluginInfoFromRepositoryForPlatformArgsForCall(0)
								Expect(pluginNameArg).To(Equal("some-plugin"))
								Expect(pluginRepositoryArg).To(Equal(configv3.PluginRepository{Name: "some-repo", URL: "http://some-url"}))
								Expect(pluginPlatform).To(Equal("linux64"))

								Expect(fakeActor.DownloadExecutableBinaryFromURLCallCount()).To(Equal(1))
								urlArg, dirArg := fakeActor.DownloadExecutableBinaryFromURLArgsForCall(0)
								Expect(urlArg).To(Equal("http://some-url"))
								Expect(dirArg).To(ContainSubstring("temp"))
							})
						})

						Context("when getting the binary succeeds", func() {
							BeforeEach(func() {
								fakeActor.DownloadExecutableBinaryFromURLReturns("some-path", 4, nil)
							})

							Context("when the checksum fails", func() {
								BeforeEach(func() {
									fakeActor.ValidateFileChecksumReturns(false)
								})

								It("returns the checksum error", func() {
									Expect(executeErr).To(MatchError(InvalidChecksumError{}))

									Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin..."))
									Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.3 found in: some-repo"))
									Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.2 is already installed."))
									Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors."))
									Expect(testUI.Out).To(Say("Install and use plugins at your own risk."))
									Expect(testUI.Out).To(Say("Starting download of plugin binary from repository some-repo..."))
									Expect(testUI.Out).To(Say("4 bytes downloaded..."))
									Expect(testUI.Out).ToNot(Say("Installing plugin"))

									Expect(fakeActor.ValidateFileChecksumCallCount()).To(Equal(1))
									pathArg, checksumArg := fakeActor.ValidateFileChecksumArgsForCall(0)
									Expect(pathArg).To(Equal("some-path"))
									Expect(checksumArg).To(Equal("some-checksum"))
								})
							})

							Context("when the checksum succeeds", func() {
								BeforeEach(func() {
									fakeActor.ValidateFileChecksumReturns(true)
								})

								Context("when creating an executable copy errors", func() {
									BeforeEach(func() {
										fakeActor.CreateExecutableCopyReturns("", errors.New("some-error"))
									})

									It("returns the error", func() {
										Expect(executeErr).To(MatchError(errors.New("some-error")))

										Expect(testUI.Out).To(Say("4 bytes downloaded..."))
										Expect(testUI.Out).ToNot(Say("Installing plugin"))

										Expect(fakeActor.CreateExecutableCopyCallCount()).To(Equal(1))
										pathArg, tempDirArg := fakeActor.CreateExecutableCopyArgsForCall(0)
										Expect(pathArg).To(Equal("some-path"))
										Expect(tempDirArg).To(ContainSubstring("temp"))
									})
								})

								Context("when creating an exectubale copy succeeds", func() {
									BeforeEach(func() {
										fakeActor.CreateExecutableCopyReturns("copy-path", nil)
									})

									Context("when validating the new plugin errors", func() {
										BeforeEach(func() {
											fakeActor.GetAndValidatePluginReturns(configv3.Plugin{}, pluginaction.PluginInvalidError{})
										})

										It("returns the error", func() {
											Expect(executeErr).To(MatchError(shared.PluginInvalidError{}))

											Expect(testUI.Out).To(Say("4 bytes downloaded..."))
											Expect(testUI.Out).ToNot(Say("Installing plugin"))

											Expect(fakeActor.GetAndValidatePluginCallCount()).To(Equal(1))
											_, commandsArg, tempDirArg := fakeActor.GetAndValidatePluginArgsForCall(0)
											Expect(commandsArg).To(Equal(Commands))
											Expect(tempDirArg).To(Equal("copy-path"))
										})
									})

									Context("when validating the new plugin succeeds", func() {
										BeforeEach(func() {
											fakeActor.GetAndValidatePluginReturns(configv3.Plugin{
												Name:    "some-plugin",
												Version: configv3.PluginVersion{Major: 1, Minor: 2, Build: 3},
											}, nil)
										})

										Context("when uninstalling the existing errors", func() {
											BeforeEach(func() {
												expectedErr = errors.New("uninstall plugin error")
												fakeActor.UninstallPluginReturns(expectedErr)
											})

											It("returns the error", func() {
												Expect(executeErr).To(MatchError(expectedErr))

												Expect(testUI.Out).To(Say("Uninstalling existing plugin..."))
												Expect(testUI.Out).ToNot(Say("Plugin some-plugin successfully uninstalled\\."))

												Expect(fakeActor.UninstallPluginCallCount()).To(Equal(1))
												_, pluginNameArg := fakeActor.UninstallPluginArgsForCall(0)
												Expect(pluginNameArg).To(Equal("some-plugin"))
											})
										})

										Context("when uninstalling the existing plugin succeeds", func() {
											Context("when installing the new plugin errors", func() {
												BeforeEach(func() {
													expectedErr = errors.New("install plugin error")
													fakeActor.InstallPluginFromPathReturns(expectedErr)
												})

												It("returns the error", func() {
													Expect(executeErr).To(MatchError(expectedErr))

													Expect(testUI.Out).To(Say("Plugin some-plugin successfully uninstalled."))
													Expect(testUI.Out).To(Say("Installing plugin some-plugin..."))
													Expect(testUI.Out).ToNot(Say("successfully installed"))

													Expect(fakeActor.InstallPluginFromPathCallCount()).To(Equal(1))
													pathArg, pluginArg := fakeActor.InstallPluginFromPathArgsForCall(0)
													Expect(pathArg).To(Equal("copy-path"))
													Expect(pluginArg).To(Equal(configv3.Plugin{
														Name:    "some-plugin",
														Version: configv3.PluginVersion{Major: 1, Minor: 2, Build: 3},
													}))
												})
											})

											Context("when installing the new plugin succeeds", func() {
												It("uninstalls the existing plugin and installs the new one", func() {
													Expect(executeErr).ToNot(HaveOccurred())

													Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin..."))
													Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.3 found in: some-repo"))
													Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.2 is already installed."))
													Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors."))
													Expect(testUI.Out).To(Say("Install and use plugins at your own risk."))
													Expect(testUI.Out).To(Say("Starting download of plugin binary from repository some-repo..."))
													Expect(testUI.Out).To(Say("4 bytes downloaded..."))
													Expect(testUI.Out).To(Say("Uninstalling existing plugin..."))
													Expect(testUI.Out).To(Say("OK"))
													Expect(testUI.Out).To(Say("Plugin some-plugin successfully uninstalled."))
													Expect(testUI.Out).To(Say("Installing plugin some-plugin..."))
													Expect(testUI.Out).To(Say("OK"))
													Expect(testUI.Out).To(Say("some-plugin 1.2.3 successfully installed"))
												})
											})
										})
									})
								})
							})
						})
					})

					Context("when the plugin is NOT already installed", func() {
						Context("when getting the binary errors", func() {
							BeforeEach(func() {
								expectedErr = errors.New("some-error")
								fakeActor.DownloadExecutableBinaryFromURLReturns("", 0, expectedErr)
							})

							It("returns the error", func() {
								Expect(executeErr).To(MatchError(expectedErr))

								Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin..."))
								Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.3 found in: some-repo"))
								Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors."))
								Expect(testUI.Out).To(Say("Install and use plugins at your own risk."))
								Expect(testUI.Out).To(Say("Starting download of plugin binary from repository some-repo..."))

								Expect(testUI.Out).ToNot(Say("downloaded"))
								Expect(fakeActor.GetAndValidatePluginCallCount()).To(Equal(0))
							})
						})

						Context("when getting the binary succeeds", func() {
							BeforeEach(func() {
								fakeActor.DownloadExecutableBinaryFromURLReturns("some-path", 4, nil)
							})

							Context("when the checksum fails", func() {
								BeforeEach(func() {
									fakeActor.ValidateFileChecksumReturns(false)
								})

								It("returns the checksum error", func() {
									Expect(executeErr).To(MatchError(InvalidChecksumError{}))

									Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin..."))
									Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.3 found in: some-repo"))
									Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors."))
									Expect(testUI.Out).To(Say("Install and use plugins at your own risk."))
									Expect(testUI.Out).To(Say("Starting download of plugin binary from repository some-repo..."))
									Expect(testUI.Out).To(Say("4 bytes downloaded..."))
									Expect(testUI.Out).ToNot(Say("Installing plugin"))
								})
							})

							Context("when the checksum succeeds", func() {
								BeforeEach(func() {
									fakeActor.ValidateFileChecksumReturns(true)
								})

								Context("when creating an executable copy errors", func() {
									BeforeEach(func() {
										fakeActor.CreateExecutableCopyReturns("", errors.New("some-error"))
									})

									It("returns the error", func() {
										Expect(executeErr).To(MatchError(errors.New("some-error")))

										Expect(testUI.Out).To(Say("4 bytes downloaded..."))
										Expect(testUI.Out).ToNot(Say("Installing plugin"))
									})
								})

								Context("when creating an executable copy succeeds", func() {
									BeforeEach(func() {
										fakeActor.CreateExecutableCopyReturns("copy-path", nil)
									})

									Context("when validating the plugin errors", func() {
										BeforeEach(func() {
											fakeActor.GetAndValidatePluginReturns(configv3.Plugin{}, pluginaction.PluginInvalidError{})
										})

										It("returns the error", func() {
											Expect(executeErr).To(MatchError(shared.PluginInvalidError{}))

											Expect(testUI.Out).To(Say("4 bytes downloaded..."))
											Expect(testUI.Out).ToNot(Say("Installing plugin"))
										})
									})

									Context("when validating the plugin succeeds", func() {
										BeforeEach(func() {
											fakeActor.GetAndValidatePluginReturns(configv3.Plugin{
												Name:    "some-plugin",
												Version: configv3.PluginVersion{Major: 1, Minor: 2, Build: 3},
											}, nil)
										})

										Context("when installing the plugin errors", func() {
											BeforeEach(func() {
												expectedErr = errors.New("install plugin error")
												fakeActor.InstallPluginFromPathReturns(expectedErr)
											})

											It("returns the error", func() {
												Expect(executeErr).To(MatchError(expectedErr))

												Expect(testUI.Out).To(Say("Installing plugin some-plugin..."))
												Expect(testUI.Out).ToNot(Say("successfully installed"))
											})
										})

										Context("when installing the plugin succeeds", func() {
											It("uninstalls the existing plugin and installs the new one", func() {
												Expect(executeErr).ToNot(HaveOccurred())

												Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin..."))
												Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.3 found in: some-repo"))
												Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors."))
												Expect(testUI.Out).To(Say("Install and use plugins at your own risk."))
												Expect(testUI.Out).To(Say("Starting download of plugin binary from repository some-repo..."))
												Expect(testUI.Out).To(Say("4 bytes downloaded..."))
												Expect(testUI.Out).To(Say("Installing plugin some-plugin..."))
												Expect(testUI.Out).To(Say("OK"))
												Expect(testUI.Out).To(Say("some-plugin 1.2.3 successfully installed"))
											})
										})
									})
								})
							})
						})
					})
				})

				Context("when the -f argument is not given (user is prompted for confirmation)", func() {
					BeforeEach(func() {
						fakeActor.ValidateFileChecksumReturns(true)
					})

					Context("when the plugin is already installed", func() {
						BeforeEach(func() {
							fakeConfig.GetPluginReturns(configv3.Plugin{
								Name:    "some-plugin",
								Version: configv3.PluginVersion{Major: 1, Minor: 2, Build: 2},
							}, true)
							fakeActor.IsPluginInstalledReturns(true)
						})

						Context("when the user chooses no", func() {
							BeforeEach(func() {
								input.Write([]byte("n\n"))
							})

							It("cancels plugin installation", func() {
								Expect(executeErr).ToNot(HaveOccurred())

								Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin\\.\\.\\."))
								Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.3 found in: some-repo"))
								Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.2 is already installed\\."))
								Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
								Expect(testUI.Out).To(Say("Install and use plugins at your own risk\\."))
								Expect(testUI.Out).To(Say("Do you want to uninstall the existing plugin and install some-plugin 1\\.2\\.3\\? \\[yN\\]"))
								Expect(testUI.Out).To(Say("Plugin installation cancelled\\."))
							})
						})

						Context("when the user chooses the default", func() {
							BeforeEach(func() {
								input.Write([]byte("\n"))
							})

							It("cancels plugin installation", func() {
								Expect(executeErr).ToNot(HaveOccurred())

								Expect(testUI.Out).To(Say("Do you want to uninstall the existing plugin and install some-plugin 1.2.3\\? \\[yN\\]"))
								Expect(testUI.Out).To(Say("Plugin installation cancelled\\."))
							})
						})

						Context("when the user input is invalid", func() {
							BeforeEach(func() {
								input.Write([]byte("e\n"))
							})

							It("returns an error", func() {
								Expect(executeErr).To(HaveOccurred())

								Expect(testUI.Out).To(Say("Do you want to uninstall the existing plugin and install some-plugin 1.2.3\\? \\[yN\\]"))
								Expect(testUI.Out).ToNot(Say("Plugin installation cancelled\\."))
								Expect(testUI.Out).ToNot(Say("Installing plugin"))
							})
						})

						Context("when the user chooses yes", func() {
							BeforeEach(func() {
								input.Write([]byte("y\n"))
								fakeActor.DownloadExecutableBinaryFromURLReturns("some-path", 4, nil)
								fakeActor.CreateExecutableCopyReturns("copy-path", nil)
								fakeActor.GetAndValidatePluginReturns(configv3.Plugin{
									Name:    "some-plugin",
									Version: configv3.PluginVersion{Major: 1, Minor: 2, Build: 3},
								}, nil)
							})

							It("installs the plugin", func() {
								Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin\\.\\.\\."))
								Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.3 found in: some-repo"))
								Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.2 is already installed\\."))
								Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
								Expect(testUI.Out).To(Say("Install and use plugins at your own risk\\."))
								Expect(testUI.Out).To(Say("Do you want to uninstall the existing plugin and install some-plugin 1\\.2\\.3\\? \\[yN\\]"))
								Expect(testUI.Out).To(Say("Starting download of plugin binary from repository some-repo..."))
								Expect(testUI.Out).To(Say("4 bytes downloaded..."))
								Expect(testUI.Out).To(Say("Uninstalling existing plugin..."))
								Expect(testUI.Out).To(Say("OK"))
								Expect(testUI.Out).To(Say("Plugin some-plugin successfully uninstalled."))
								Expect(testUI.Out).To(Say("Installing plugin some-plugin..."))
								Expect(testUI.Out).To(Say("OK"))
								Expect(testUI.Out).To(Say("some-plugin 1.2.3 successfully installed"))
							})
						})
					})

					Context("when the plugin is NOT already installed", func() {
						Context("when the user chooses no", func() {
							BeforeEach(func() {
								input.Write([]byte("n\n"))
							})

							It("cancels plugin installation", func() {
								Expect(executeErr).ToNot(HaveOccurred())

								Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin\\.\\.\\."))
								Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.3 found in: some-repo"))
								Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
								Expect(testUI.Out).To(Say("Install and use plugins at your own risk\\."))
								Expect(testUI.Out).To(Say("Do you want to install the plugin some-plugin\\? \\[yN\\]"))
								Expect(testUI.Out).To(Say("Plugin installation cancelled\\."))
							})
						})

						Context("when the user chooses the default", func() {
							BeforeEach(func() {
								input.Write([]byte("\n"))
							})

							It("cancels plugin installation", func() {
								Expect(executeErr).ToNot(HaveOccurred())

								Expect(testUI.Out).To(Say("Do you want to install the plugin some-plugin\\? \\[yN\\]"))
								Expect(testUI.Out).To(Say("Plugin installation cancelled\\."))
							})
						})

						Context("when the user input is invalid", func() {
							BeforeEach(func() {
								input.Write([]byte("e\n"))
							})

							It("returns an error", func() {
								Expect(executeErr).To(HaveOccurred())

								Expect(testUI.Out).To(Say("Do you want to install the plugin some-plugin\\? \\[yN\\]"))
								Expect(testUI.Out).ToNot(Say("Plugin installation cancelled\\."))
								Expect(testUI.Out).ToNot(Say("Installing plugin"))
							})
						})

						Context("when the user chooses yes", func() {
							BeforeEach(func() {
								input.Write([]byte("y\n"))
								fakeActor.DownloadExecutableBinaryFromURLReturns("some-path", 4, nil)
								fakeActor.CreateExecutableCopyReturns("copy-path", nil)
								fakeActor.GetAndValidatePluginReturns(configv3.Plugin{
									Name:    "some-plugin",
									Version: configv3.PluginVersion{Major: 1, Minor: 2, Build: 3},
								}, nil)
							})

							It("installs the plugin", func() {
								Expect(executeErr).ToNot(HaveOccurred())

								Expect(testUI.Out).To(Say("Searching some-repo for plugin some-plugin\\.\\.\\."))
								Expect(testUI.Out).To(Say("Plugin some-plugin 1\\.2\\.3 found in: some-repo"))
								Expect(testUI.Out).To(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
								Expect(testUI.Out).To(Say("Install and use plugins at your own risk\\."))
								Expect(testUI.Out).To(Say("Do you want to install the plugin some-plugin\\? \\[yN\\]"))
								Expect(testUI.Out).To(Say("Starting download of plugin binary from repository some-repo..."))
								Expect(testUI.Out).To(Say("4 bytes downloaded..."))
								Expect(testUI.Out).To(Say("Installing plugin some-plugin..."))
								Expect(testUI.Out).To(Say("OK"))
								Expect(testUI.Out).To(Say("some-plugin 1.2.3 successfully installed"))
							})
						})
					})
				})
			})
		})
	})
})
