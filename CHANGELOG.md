# Changelog

## v3.1.1 (2021-04-29)

- Add go.mod
- Maintenance

## v3.1.0 (2018-09-04)

- Fix flag settings scope
- Fix command flag noempty tag
- Implement flag handlers
- Add Cmd.Options.Logger
- Add Cmd.Options.ExitOnError
- Add Cmd.Options.ConfigType

## v3.0.2 (2018-04-01)

- Fix slice flag value issue

## v3.0.1 (2018-04-01)

- Fix field index issue
- Improve tests and code coverage
- Add Cmd.exit method

## v3.0.0 (2018-02-11)

- [BREAKING CHANGE] Commands require struct tag
- [BREAKING CHANGE] Unknown arguments cause error as default
- Implement unknown argument handling
- Implement noempty arguments
- Add Cmd.FlagArgs
- Add Cmd.Options.AnyError
- Improve tests and examples

## v2.2.0 (2018-02-02)

- Improve error handling for required flags
- Add env value to usage content
- Implement global arguments
- Add Flag.FormattedArg method

## v2.1.0 (2018-01-31)

- Implement auto usage and version printing
- Add Cmd.FlagValue method
- Add FlagSet.FlagByArg method
- Add Flag.Value method

## v2.0.0 (2018-01-15)

- [BREAKING CHANGE] Change return values for New table function
- [BREAKING CHANGE] Drop row number argument from Table.AddRow method
- Add Table.SetRow method
- Improve Cmd.usageContent
- Fix flag arguments issue

## v1.0.1 (2018-01-09)

- Fix usage issue
- Fix indexTo issue

## v1.0.0 (2018-01-07)

- Initial release
