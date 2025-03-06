using Grpc.Core;
using System.Threading.Tasks;
using MyGreeterCsharp.Api.V1;
using Serilog;
using AKSMiddleware;
using Google.Protobuf.WellKnownTypes;
using Azure.Identity;
using Azure;
using Azure.Core;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

namespace MyGreeterCsharp.Server.Handler;

public partial class MyGreeterCsharpServer
{
    public override async Task<Empty> DeleteStorageAccount(DeleteStorageAccountRequest request, ServerCallContext context)
    {
        // TODO: Implement DeleteStorageAccount
        var response = new Empty();

        return await Task.FromResult(response);
    }
}