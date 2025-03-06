# Stage 1: Build the .NET application
FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
WORKDIR /app

# Install the Azure Artifacts Credential Provider
RUN curl -L https://raw.githubusercontent.com/Microsoft/artifacts-credprovider/master/helpers/installcredprovider.sh  | sh

# Copy csproj and restore as distinct layers using nuget.config for Azure Artifacts
COPY *.csproj .
COPY ./nuget.config .
ARG PAT
ENV VSS_NUGET_EXTERNAL_FEED_ENDPOINTS='{"endpointCredentials": [{"endpoint":"https://pkgs.dev.azure.com/service-hub-flg/service_hub_validation/_packaging/service_hub_validation__PublicPackages/nuget/v3/index.json", "username":"docker", "password":"'${PAT}'"}]}'


RUN dotnet restore

# Copy the source code and build for Client and Server
COPY Src ./Src

# Publish the Client project
RUN dotnet publish -c Client -o /app/client

# Publish the Server project
RUN dotnet publish -c Server -o /app/server

# Stage 2: Create a runtime image
FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS runtime
WORKDIR /app

# Copy the published Client and Server binaries
COPY --from=build /app/client ./client
COPY --from=build /app/server ./server
