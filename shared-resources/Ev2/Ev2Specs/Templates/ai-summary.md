# AI-Summary
## Directory Summary
This directory contains templates for managing Azure resources and role assignments in the Ev2 pipeline. It includes a Bicep template for pushing images to an Azure Container Registry with a managed identity, and an ARM template for assigning roles at the subscription level.

**Tags:** Azure, Role Assignment, Ev2, Templates

## File Details
    
### /shared-resources/Ev2/Ev2Specs/Templates/AcrPush.SharedResources.Template.bicep
This Bicep template is used in the Ev2 pipeline to manage Azure resources for pushing images to an Azure Container Registry (ACR). It defines parameters for resource name and location, and creates a user-assigned managed identity for the pipeline. The identity is granted a role assignment to push images to the ACR.

### /shared-resources/Ev2/Ev2Specs/Templates/RoleAssignment.Subscription.Template.json
This is an Azure Resource Manager (ARM) template for assigning roles to a principal at the subscription level. It defines parameters for the principal ID and the built-in role type, and uses these to create a role assignment resource. The template includes predefined role definitions for roles such as Owner, Contributor, and Reader.
