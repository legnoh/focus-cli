# focus-cli

Get macOS Focus mode via CLI

## install

```sh
brew install legnoh/etc/mac-focus-cli
```

> [!IMPORTANT]
> You must set "Full Disk Access" to `focus` command.
> Because this application accesses below files.
> 
> - `~/Library/DoNotDisturb/DB/ModeConfigurations.json`
> - `~/Library/DoNotDisturb/DB/Assertions.json`
> 
> <img width="715" alt="full disk access in settings" src="https://github.com/user-attachments/assets/18f5541a-7543-47da-83f6-eadec6701bb1">
> 
> - FYI: [Controlling app access to files in macOS - Apple Support](https://support.apple.com/guide/security/controlling-app-access-to-files-secddd1d86a6/)


## usage

```sh
## get now focus mode name
focus get

## get by json format
focus get --json
```
