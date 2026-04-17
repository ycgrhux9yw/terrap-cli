# Terrap, by Sirrend
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  ![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/sirrend/terrap-cli?filename=go.mod)"/>
 required changes. </br>
The tool offers clear and actionable notifications, helping</br></br>
🔍 version project, therefore some data might be partial.**

> **Personal note:** I'm using this tool to manage provider upgrades across several AWS-heavy workspaces. Forked for local experimentation and learning.

## Resources
* Documentation - <a href="https://www.sirrend.com/terrap-docs">sirrend.com/terrap-docs</a>

## Constraints 🧱
1. Supported Terraform Core versions: `>=0.13`.
2. Every provider which uses `Terraform Core 0.13` or higher.

## Good To Know 💡
Terrap decides which Terraform version to use in the following order:
1. The latest installed Terraform version found locally.
2. If the `TERRAP_TERRAFORM_VERSION` environment variable is set, Terrap will use the version specified as long as it matches the `>=0.13` constraint.</br>
    Set environment variable on mac/linux:</br>
    ```shell
   export <var>=<value>
   ```
   Set environment variable on windows:</br>
    ```shell
   $Env:TERRAP_TERRAFORM_VERSION = "0.13"
   ```
   
4. If none of the above is applicable, Terrap will download the latest available version.

## How to Download ⬇️
### Clone sirrend/terrap-cli
```shell
git clone https://github.com/sirrend/terrap-cli
cd terrap-cli

go build -o terrap .

chmod +x terrap
mv terrap /usr/local/bin/
```

### Brew
```shell
brew tap sirrend/products
brew install terrap
```

Validate terrap is working by executing `terrap`.

## Quick Start ⏩

### Initialize my First Workspace
1. `CD` to the local Terraform folder you want to work with.</br>
   `cd < /terraform/folder/path >`</br></br>

2terrap init`.</br></br>
    <strong>Important!</strong> </br>
    As Terrap runs <code>terraform init</code> under the hood, it would need every configuration component you normally use when executing the command</br>
    It can be environment variables, the <code>.aws/n

3. Scan your workspace with: `terrap scan`

https://user-images.githubusercontent044850-3473952a-4169-4d63-beb7-cf1664afc35a.mov

## Features 🚀
### Scan for changes with `scan`
Scan your infrastructure for changes in the following provider version for a safe and easy upgrade!</br>
Looking for a specific resource type changes? Use the `--data-sources` `--resources` and `--provider` flags.

### Stay up-to-date with `whats-new`
Ready to explore what's new in the following version of your provider? Simply execute `terrap whats-new`.</br>
Looking to delve into a specific version of your

## Local Dev Notes 🗒️
- When testing against real AWS workspaces, make sure `AWS_PROFILE` is set before running `terrap scan`.
- To avoid accidental state changes, always run `terrap scan` in a read-only context — never from a directory with active `terraform.tfstate` locks.
- Useful alias: `alias tscan='AWS_PROFILE=dev terrap scan'`
