# Ev2 vs Normal Rollout

We currently have two processes to roll out the service.

## [Ev2](https://ev2docs.azure.net/getting-started/overview.html)

- For Test, Production (Azure public cloud), sovereign, and air-gapped clouds. More details about [Ev2 Cloud terminology](https://ev2docs.azure.net/getting-started/production/release-pipeline.html#what-is-new-in-governed-yaml-pipelines-with-ev2-managed-sdp).
- Uses a set of json config files (named artifacts) to orchestrate the deployment of your service.

## Normal Rollout

- For developer to test if the microservices can work together quickly.
- We provide a set of Makefile targets to orchestrate the deployment of your service. It doesn't use Ev2.

Ev2 rollout and normal rollout largely share the same files, so you will typically only have to make changes to your files once. There are a few notable differences described below.

## Files Shared by Ev2 and Normal Rollout

- all bicep files
- `template-ServiceResources.Parameters.json`
- `template-Main.SharedResources.Parameters.json`
- all files under the `deployments` directory of each service

## Files **NOT** Shared by Ev2 and Normal Rollout

There are several files that exist specifically for Ev2 buildout/rollout.

### Ev2 Specific

All files that are in the Ev2 directories. We have an Ev2 directory in shared-resources and service directory.

## Replacing Templated Values: Using `template-` files in Ev2

Files that contain [Go template "actions"](https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go#step-4-writing-a-template) need to be executed (replacing the placeholder with real value). They have the prefix `template-` in their file name. Examples include:

- `template-ServiceResources.Parameters.json`
- `template-Main.SharedResources.Parameters.json`
- `template-common-values.yaml`

The Ev2 rollout replaces these placeholders through the `ScopeBinding.json` file, while the normal rollout replaces these placeholders using the `env-config.yaml` file. The `env-config.yaml` is generated from the Makefile target in the README directory.

# Ev2 Setup

Make sure you have followed the necessary steps for Ev2 mentioned in the [Service Hub README.md](https://dev.azure.com/service-hub-flg/service_hub/_git/service_hub?path=/README.md&_a=preview).

## Setting up subscriptions for Ev2

### Using a pre-existing subscription instead of provisioning through Ev2

- If you want to use an existing subscription for all regions instead of provisioning a subscription for each release region, you will need to perform the following steps.
  - Remove `SubscriptionProvisioning.Parameters.json` from shared-resources\Ev2\Ev2Specs\Parameters
  - Remove the following code from `serviceModel.json` in shared-resources\Ev2\Ev2Specs

    ```json
      "subscriptionProvisioning": {
        "rolloutParametersPath": "Parameters\\SubscriptionProvisioning.Parameters.json",
        "scopeTags": [
          {
            "name": "sharedInputs"
          },
          {
            "name": "subscriptionInputs"
          }
        ]
      },
    ```

  - Follow the [steps](https://ev2docs.azure.net/features/service-artifacts/actions/subscriptionProvisioningParameters.html#subscription-registrationbackfill) to register your subscription as declarative  backfill.
  - Change how you define "subscriptionKey" in all your serviceModel.json files to match the key you registered your subscription with. Including the service specific models.
  - Follow the [instructions](https://ev2docs.azure.net/features/rollout-infra/prod.html) to give your security group team permissions to your subscription.

## What do the Ev2 files provide?

### Common through all directories

1. **Configurations** - Defines the tenant id/subscription key for the Ev2 rollout and rollout environment for each of the two options below.
    1. **Production**
    2. **Test**
2. **Templates** - Defines the resources that will be rolled out in accordance to ServiceModel.json
3. **Parameters** - Defines the parameters files used by resources defined in Templates and ServiceModel.json
4. **RolloutSpec.json** - Defines what resources to deploy from the service model and the order for deploying those resources
5. **ScopeBinding.json** - Scope bindings enable you to reuse your parameters across actions. It defines resource parameters for different environments without creating different files for every environment.
6. **ServiceModel.json** - Defines the Azure subscription, resource groups, and ARM templates.
7. **Version.txt** - The file that stores version of the Ev2 build. Written to by the build script.

### **shared-resources** specific

1. Within **Templates**
    1. `AcrPush.SharedResources.Template.bicep`: Creates the identity and role assignment that allows for the publishing of the image to the ACR in our Ev2 pipeline. All microservices use these resources to publish the docker images to the ACR. In other words, these resources are *created* in shared resources, but are *used* by the services.
    2. `RoleAssignment.Subscription.Template.json`: Defines a variety of role assignments that get tied to the subscription we provision for the region we release in. The corresponding parameters file defines the exact role assignments that get created.
2. Within **Parameters**
    1. `AcrPush.SharedResources.Parameters.json`: Parameters file for `AcrPush.Sharedresources.Template.bicep`.
    2. `RoleAssignment.Subscription.Parameters.json`: Parameters file for `RoleAssignment.Subscription.Template.json`.
    3. `SubscriptionProvisioning.Parameters.json`: Parameters file used for provisioning the region based subscription defined in ServiceModel.json.
    4. `RegisterResourceProvider.Parameters.json`: Parameters file used for registering resource providers, specifically Microsoft.Compute such that clusters are able to deploy successfully.
    5. `Kubectl.Rollout.Parameters.json`: Defines the process to cluster for kubectl

### Service specific

1. **Templates**: This directory does not exist in the microservice directories. It is added in by the build pipeline during the build process. The build pipeline takes bicep files stored in serviceDirectoryName/resources and converts them to arm json files.
2. Within **Parameters**
    1. `Helm.Rollout.Parameters.json`: Defines the process to cluster for helm
    2. `PublishImage.Parameters.json`: Defines the process that pushes service image to acr.
3. **Shell**
    1. `push-image-to-acr`: The script used to push the service image to the acr.

## Build and Release with Ev2 through OneBranch Pipelines

We have created a generalized [build script](../pipeline-files/buildEv2Artifacts.sh) that can build both shared-resources and any service. The script gets called by the build pipeline to generate the required artifacts for an Ev2 Deployment.

### (Optional) Testing Builds Locally

The build script can be also be run locally if you want to see what artifacts will get published by the build pipeline.

Below are the required arguments the build script gets run with.

| Variable    | Purpose   |  Options |
|---|---|---|
|directoryName| The name of the directory you are performing the build for|e.g shared-resources|
|outputDir|The **full** path of the output directory you want the artifacts to be stored in |---|
|isService|Is the build for a service or for shared-resources |true/false|
|rolloutInfra|Which configuration file gets used when building. |Prod/Test|
|buildNumber| The version that gets associated with this set of artifacts. However you decide to version local artifacts |e.g. "1.0.0" or "20240820" |
|isLocal| Are you building locally? | true/false |

For example, building artifacts for Ev2 for the shared-resources directory, can look like this.

```bash
./pipeline-files/buildEv2Artifacts.sh shared-resources ~/shared-artifacts false Prod "1.0.0" true
```

### (Optional) Adding customizations to builds

We understand that not every service will share the exact same minimum build that we have created, and might require additional files to be a part of the published artifacts. As a result, you can add a script under the **server** directory of your service with the name "buildCustomEv2.sh", and our build will automatically call it if the file exists.

### Creating OneBranch Build/Release Pipelines

1. The most basic build and release pipeline yaml files for shared and service specific resources have been provided. The build pipeline exists under **pipeline-files** and the release pipeline files exist under their corresponding generated directories. Below is a deeper description of what we provide.

| FileName   | Type   |  Purpose | Can be re-used? | Required Variables |
|---|---|---|---|---|
| OneBranch.Official.Build | Build | Build yaml file that builds service image, and calls a script that tests service, packages helm deployments files, converts bicep to arm json files, and compiles all files into an "artifact" directory. |Yes! As long as pipeline variables are changed as required and service directory holds expected structure.| <ul><li>directoryName: The name of the directory where the service/shared-resources code is located.</li><li>isService: true if the pipeline is for a service, false if it is for shared-resources.</li><li>rolloutInfra: Prod or Test based on which configuration file needs to be used</li><li>forceReinstallCredentialProvider: true (to avoid credential provider caching issues)</li></ul> |
| OneBranch.Official.Release | Release | Release yaml file that links to the necessary service build pipeline, extracts the service artifacts from the build, and releases in the mentioned environments and regions. | Semi-reusable. Each service will have a copy of the release yaml file with their corresponding build pipeline defined. The source of the build pipeline cannot be a variable taken in at runtime, so it must be hardcoded.|N/A|

2. For each Build and for each Release you will need to create a seperate OneBranch yaml pipeline. The recommended method can be found in these [instructions](https://eng.ms/docs/products/onebranch/onboarding/onebranchresources/newonebranchpipeline). Make sure to select existing yaml pipeline and select the corresponding yaml file stored. For builds they will be in **pipeline-files** and for releases they will be in **shared-resources** and service directories.

3. Once the pipelines are built, add the required variables mentioned in the above table.

### Releasing in certain regions or all regions

Currently, the release pipeline yaml files have listed 4 regions that will be released in. If you want to release to more regions, you must include the region names in the following line of your release file.

```
Select: regions(australiaeast, eastus2, swedencentral, southeastasia)
```

If you do not want to manually list all your release regions, you can use Ev2's concept of service presence. This will target all the regions that have presence in, so please make sure the regions you want to target have presence registered. If you are selecting specific regions, presence is not required to be registered.

- To register your service presence follow these [instructions](https://ev2docs.azure.net/getting-started/production/presence.html?q=service%20presence&tabs=powershell#one-time-registration-of-service-presence-for-existing-regions)
- Once you have registered your service in the regions, you can change the selection of regions in your release pipeline to the following line. This will release in **ALL** regions your service presence was registered in.

  ```
  Select: regions(*)
  ```

### Orchestrating Build and Release

1. Run the build pipeline. Make sure the variables

   The following artifacts will be published by the Ev2 Build

- For shared-resources: drop_combineArtifacts_prepare
  - Ev2Specs
    - Build
      - argo-rollouts-namespace.yaml
      - install.yaml
    - Configuration.json
    - Configurations
      - Prod
        - Configuration.json
      - Test
        - Configuration.json
    - Parameters
      - AcrPush.SharedResources.Parameters.json
      - Main.SharedResources.Parameters.json
      - RoleAssignment.Subscription.Parameters.json
      - SubscriptionProvisioning.Parameters.json
      - RegisterResourceProvider.Parameters.json
    - RolloutSpec.json
    - ScopeBinding.json
    - ServiceModel.json
    - Templates
      - AcrPush.SharedResources.Template.bicep
      - AcrPush.SharedResources.Template.json
      - Main.SharedResources.Template.bicep
      - Main.SharedResources.Template.json
      - RoleAssignment.Subscription.Template.json
    - Version.txt
- For service specific: drop_combineArtifacts_prepare
  - Ev2Specs
    - Build
      - serviceDirectoryName-0.1.0.tgz
      - serviceDirectoryName-image-metadata.json
      - serviceDirectoryName-image.tar
      - values-client.yaml
      - values-common.yaml
      - values-demoserver.yaml
      - values-server.yaml
    - Configuration.json
    - Configurations
      - Prod
        - Configuration.json
      - Test
        - Configuration.json
    - Parameters
      - ServiceResources.Parameters.json
      - Helm.Rollout.Parameters.json
      - PublishImage.Parameters.json
    - RolloutSpec.json
    - ScopeBinding.json
    - ServiceModel.json
    - Shell
      - LICENSE
      - README.md
      - crane
      - crane.tar.gz
      - gcrane
      - push-image-to-acr.sh
    - Templates
      - Azuresdk.ServiceResources.Template.bicep
      - Azuresdk.ServiceResources.Template.json
    - Version.txt
    - push-image-to-acr.tar

2. Run the release pipeline
3. Once the release pipeline has completed its "PreReleaseValidation" job, it will start it's "ApprovalService" job. The pipeline runner will receive an email with a link to the [approval service page](https://approval.azengsys.com/Home/PendingRelease). A person on the release approvers team (that is not the submitter) must approve the release before the pipeline can go any further.
4. Monitor Ev2 rollout by using the link provided in the release pipeline's "PROD_Managed_SDP_Monitoring" job's logs.

# Additional Information

## Helpful Ev2 Docs

- [Getting Started](https://ev2docs.azure.net/getting-started/overview.html)
- [Overview of Service Artifacts](https://ev2docs.azure.net/features/service-artifacts/service_artifacts_overview.html)
- [Actions in Ev2 Explained for Rollout Orchestration](https://ev2docs.azure.net/features/rollout-orchestration/actions.html?q=actions)
- [Actions in Ev2 for Service Artifacts](https://ev2docs.azure.net/features/service-artifacts/actions/overview.html) Includes kubernetes application deployment action.
- [Helm Application Deployment to Kubernetes in Ev2](https://ev2docs.azure.net/features/service-artifacts/actions/kubernetes-app/helm/app-modeling.html?q=helm%20appli&tabs=tabid-1)
- [Shell Extension Artifacts](https://ev2docs.azure.net/features/service-artifacts/actions/shell-extensions/artifacts.html)
- [Shell Extension Overview](https://ev2docs.azure.net/features/service-artifacts/actions/shell-extensions/overview.html?q=shell%20extension)
- [Publish Image to ACR in Ev2](https://eng.ms/docs/cloud-ai-platform/azure-edge-platform-aep/aep-engineering-systems/productivity-and-experiences/ce-legacy-infrastructure/onebranch/build/containerbasedworkflow/dockerimagesandacr/publishdockerimagesusingev2)
- [Store Helm Charts](https://learn.microsoft.com/en-us/azure/container-registry/container-registry-helm-repos)

## Deploying Service Components

We currently have our service artifacts set up such that we can deploy our service components (client, server, and demoserver) in parallel. Should you want to deploy your components sequentially, follow the sequential deployment example below. You will need to change the SRDef in your ServiceModel.json.

### Parallel Deployment (Current)

Separate out the different components of the service such that they can deploy in parallel.

RolloutSpec.json

```json
    {
      "name": "HelmDeploy-serviceDirectoryName-server",
      "targetType": "ApplicationDefinition",
      "applications": {
        "names": [
          "serviceDirectoryNameserver"
        ],
        "actions": [
          "AppDeploy"
        ],
        "applyAcrossServiceResources": {
          "definitionName": "serviceDirectoryNameserver-SRDef"
        }
      },
      "dependsOn": [
        "DeployServiceResources",
        "PublishImageToAcr"
      ]
    },
    {
      "name": "HelmDeploy-serviceDirectoryName-demoserver",
      "targetType": "ApplicationDefinition",
      "applications": {
        "names": [
          "serviceDirectoryNamedemoserver"
        ],
        "actions": [
          "AppDeploy"
        ],
        "applyAcrossServiceResources": {
          "definitionName": "serviceDirectoryNamedemoserver-SRDef"
        }
      },
      "dependsOn": [
        "DeployServiceResources",
        "PublishImageToAcr"
      ]
    }
```

ServiceModel.json

```json
        {
          "name": "serviceDirectoryNameclient-SRDef",
          "composedOf": {
            "application": {
              "names": [
                "serviceDirectoryNameclient"
              ]
            },
            "extension": {
              "rolloutParametersPath": "Parameters\\Helm.Rollout.Parameters.json"
            }
          },
          "scopeTags": [
            {
              "name": "sharedInputs"
            },
            {
              "name": "HelmInputs"
            }
          ]
        },
        {
          "name": "serviceDirectoryNameserver-SRDef",
          "composedOf": {
            "application": {
              "names": [
                "serviceDirectoryNameserver"
              ]
            },
            "extension": {
              "rolloutParametersPath": "Parameters\\Helm.Rollout.Parameters.json"
            }
          },
          "scopeTags": [
            {
              "name": "sharedInputs"
            },
            {
              "name": "HelmInputs"
            }
          ]
        },
        {
          "name": "serviceDirectoryNamedemoserver-SRDef",
          "composedOf": {
            "application": {
              "names": [
                "serviceDirectoryNamedemoserver"
              ]
            },
            "extension": {
              "rolloutParametersPath": "Parameters\\Helm.Rollout.Parameters.json"
            }
          },
          "scopeTags": [
            {
              "name": "sharedInputs"
            },
            {
              "name": "HelmInputs"
            }
          ]
        },
```

### Sequential Deployment

RolloutSpec.json

```json
    {
      "name": "HelmDeploy-serviceDirectoryName",
      "targetType": "ApplicationDefinition",
      "applications": {
        "names": [
            "serviceDirectoryNameclient",
            "serviceDirectoryNameserver",
            "serviceDirectoryNamedemoserver"
        ],
        "actions": [
          "AppDeploy"
        ],
        "applyAcrossServiceResources": {
          "definitionName": "serviceDirectoryName-SRDef"
        }
      },
      "dependsOn": [
        "DeployServiceResources",
        "PublishImageToAcr"
      ]
    }
```

ServiceModel.json

```json
        {
          "name": "serviceDirectoryName-SRDef",
          "composedOf": {
            "application": {
              "names": [
                "serviceDirectoryNameclient",
                "serviceDirectoryNameserver",
                "serviceDirectoryNamedemoserver"
              ]
            },
            "extension": {
              "rolloutParametersPath": "Parameters\\Helm.Rollout.Parameters.json"
            }
          },
          "scopeTags": [
            {
              "name": "sharedInputs"
            },
            {
              "name": "HelmInputs"
            }
          ]
        }
```
