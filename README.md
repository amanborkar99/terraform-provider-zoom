# terraform-provider-zoom
# ZOOM PROVIDER 
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) ![Open Source? Yes!](https://badgen.net/badge/Open%20Source%20%3F/Yes%21/blue?icon=github)

## Prerequisites:
   * the [Terraform 0.14+ CLI](https://learn.hashicorp.com/tutorials/terraform/install-cli) installed locally. The Terraform binaries are located [here](https://releases.hashicorp.com/terraform/).

   * Build Binary file from source code from [here](https://github.com/amanborkar99/terraform-provider-zoom) OR can directly download the binary file from [here](https://drive.google.com/file/d/1PZDaZSs5B6Ch67fR-X4eIr28ttiPdAUL/view?usp=sharing). 


## After Cloning Repository:
   *  After cloning run the command in sequence :-
      ``` GO
      go mod init terraform-provider-zoom
      ```
      ``` GO
      go mod tidy
      ```
      ``` GO
      go mod vendor
      ```
    

## Running Terraform:
   * `Cd examples`.
   * Edit the terraform-configuration file, create resources or read.
   * Run `terraform init`, `terraform plan` and `terraform apply`.
      
  
