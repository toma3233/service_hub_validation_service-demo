# AI-Summary
## Directory Summary
This directory contains JSON parameter files for managing Azure resources within the Ev2 system. It includes configurations for registering resource providers, provisioning subscriptions, assigning roles, and deploying shared resources. The files use placeholders for dynamic values and reference additional templates and parameter files for role assignments and deployments.

**Tags:** Azure, JSON, parameters, provisioning, role assignment

## File Details
    
### /shared-resources/Ev2/Ev2Specs/Parameters/RegisterResourceProvider.Parameters.json
This JSON file defines parameters for registering a resource provider in Azure using an extension. It specifies the schema version, the extension name, type, version, and connection properties such as maximum execution time and authentication type. It also includes payload properties like waiting until completion and the namespaces of the resource provider.

### /shared-resources/Ev2/Ev2Specs/Parameters/SubscriptionProvisioning.Parameters.json
This JSON file is a parameter specification for subscription provisioning in the Ev2 system. It includes details such as subscription name, display name, workload type, billing information, and role assignment paths. The billing section uses placeholders for dynamic values, and role assignments reference specific template and parameter files.

### /shared-resources/Ev2/Ev2Specs/Parameters/RoleAssignment.Subscription.Parameters.json
This is a JSON parameters file for Azure role assignment, specifying the principal ID and built-in role type for a subscription deployment. It uses placeholders for dynamic values.

### /shared-resources/Ev2/Ev2Specs/Parameters/AcrPush.SharedResources.Parameters.json
This JSON file is an Azure deployment parameters file for shared resources, specifying the parameters 'resourcesName' and 'location' with placeholder values.
