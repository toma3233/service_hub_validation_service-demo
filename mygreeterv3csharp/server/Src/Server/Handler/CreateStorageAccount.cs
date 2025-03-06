using Grpc.Core;
using System.Threading.Tasks;
using Serilog;
using AKSMiddleware;
using Google.Protobuf.WellKnownTypes;
using Azure.Identity;
using Azure;
using Azure.Core;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
using MyGreeterCsharp.Api.V1;

namespace MyGreeterCsharp.Server.Handler;

public partial class MyGreeterCsharpServer
{
    public override async Task<CreateStorageAccountResponse> CreateStorageAccount(CreateStorageAccountRequest request, ServerCallContext context)
    {
        // TODO: implement CreateStorageAccount
        var response = new CreateStorageAccountResponse
        {
            Name = "StorageName"
        };

        return await Task.FromResult(response);
    }
}