#if CLIENT

using System.Threading.Tasks;
using System.CommandLine;

using Grpc.Net.Client;
using MyGreeterCsharp.Api.V1;
using MyGreeterCsharp.Client.Operation;

namespace MyGreeterCsharp.Client;

public static class ClientMainCommand
{
    public static async Task<int> Main(string[] args)
    {
        var rootCommand = new RootCommand("This sample service demonstrates client-server communication using gRPC and shows how to access and interact with the Azure SDK");
        rootCommand.AddCommand(StartCommand.Init());
        return await rootCommand.InvokeAsync(args);
    }
}

#endif