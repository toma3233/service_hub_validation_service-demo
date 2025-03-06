targetScope = 'subscription'

@sys.description('The name for the resources.')
param resourcesName string

@sys.description('The subscription the resources are deployed to.')
param subscriptionId string

@sys.description('The location of the resource group the resources are deployed to.')
param location string

@sys.description('The name of the resource group the resources are deployed to.')
param resourceGroupName string

module rg 'br/public:avm/res/resources/resource-group:0.2.3' = {
  name: '${resourceGroupName}Deploy'
  scope: subscription(subscriptionId)
  params: {
    name: resourceGroupName
    location: location
  }
}

module aks 'br/public:avm/res/container-service/managed-cluster:0.8.1' = {
  name: 'servicehubval-${resourcesName}-shared-resources-clusterDeploy'
  scope: resourceGroup(subscriptionId, resourceGroupName)

  params: {
    // Required parameters
    name: 'servicehubval-${resourcesName}-cluster'
    location: rg.outputs.location
    dnsPrefix: resourcesName
    primaryAgentPoolProfiles: [
      {
        name: 'agentpool'
        count: 3 // agentCount
        vmSize: 'Standard_DS11_v2'
        osType: 'Linux'
        mode: 'System'
        availabilityZones: [] // use this when availability zones ar not availabile in region
      }
    ]
    disableLocalAccounts: false
    managedIdentities: {
      systemAssigned: true
    }
    publicNetworkAccess: 'Enabled'
    omsAgentEnabled: true
    monitoringWorkspaceResourceId: workspace.outputs.resourceId
    omsAgentUseAADAuth: true
    enableOidcIssuerProfile: true
    enableWorkloadIdentity: true
    istioServiceMeshEnabled: true
    istioServiceMeshRevisions: ['asm-1-23']
  }
}

module dataCollectionRuleAssociation 'br:servicehubregistry.azurecr.io/bicep/modules/data-collection-rule-association:v5' = {
  name: 'servicehub-${resourcesName}-shared-resources-dcr-associationDeploy'
  scope: resourceGroup(subscriptionId, resourceGroupName)
  params: {
    dataCollectionRuleId: dataCollectionRule.outputs.resourceId
    aksName: aks.outputs.name
  }
}

// TODO: potentially make unique to cloud
module acr 'br/public:avm/res/container-registry/registry:0.1.1' = {
  name: 'servicehubval-${resourcesName}-${location}acrDeploy'
  scope: resourceGroup(subscriptionId, resourceGroupName)
  params: {
    name: 'servicehubval${resourcesName}${location}acr'
    location: rg.outputs.location
    roleAssignments: [
      {
        principalId: aks.outputs.kubeletIdentityObjectId
        principalType: 'ServicePrincipal'
        roleDefinitionIdOrName: 'AcrPull'
      }
    ]
  }
}

module workspace 'br/public:avm/res/operational-insights/workspace:0.3.4' = {
  name: 'servicehubval-${resourcesName}-workspaceDeploy'
  scope: resourceGroup(subscriptionId, resourceGroupName)
  params: {
    name: 'servicehubval-${resourcesName}-workspace'
    location: rg.outputs.location
  }
}

var streams = ['Microsoft-ContainerLogV2']
module dataCollectionRule 'br/public:avm/res/insights/data-collection-rule:0.1.2' = {
  name: 'servicehubval-${resourcesName}-data-collection-ruleDeploy'
  scope: resourceGroup(subscriptionId, resourceGroupName)
  params: {
    name: 'servicehubval-${resourcesName}-data-collection-rule'
    location: rg.outputs.location
    dataFlows: [
      {
        streams: streams
        destinations: [
          'ciworkspace'
        ]
      }
    ]
    dataSources: {
      extensions: [
        {
          name: 'ContainerInsightsExtension'
          streams: streams
          extensionSettings: {
            dataCollectionSettings: {
              enableContainerLogV2: true
              interval: '1m'
              namespaceFilteringMode: 'Exclude'
            }
          }
          extensionName: 'ContainerInsights'
        }
      ]
    }
    destinations: {
      logAnalytics: [
        {
          workspaceResourceId: workspace.outputs.resourceId
          name: 'ciworkspace'
        }
      ]
    }
  }
}

module serviceBusNamespace 'br/public:avm/res/service-bus/namespace:0.9.0' = {
  name: 'servicehubval-${resourcesName}-${location}-sb-nsDeploy'
  scope: resourceGroup(subscriptionId, resourceGroupName)
  params: {
    name: 'servicehubval-${resourcesName}-${location}-sb-ns'
    location: rg.outputs.location
    queues: [
      {
        name: 'servicehubval-${resourcesName}-queue'
      }
    ]
    skuObject: {
      name: 'Basic'
    }
    zoneRedundant: false
  }
}
