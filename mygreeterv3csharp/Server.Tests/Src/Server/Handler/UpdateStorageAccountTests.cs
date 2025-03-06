#if SERVER

using Moq;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using System.Threading.Tasks;
using Serilog;
using MyGreeterCsharp.Api.V1;
using MyGreeterCsharp.Server.Handler;

namespace Server.Tests;

public class UpdateStorageAccountTests
{
    private static readonly string SubscriptionId = "test-subscription";
    private static readonly string ResourceGroupName = "test-rg";
    private static readonly string ServiceAccountName = "test-service-account";

    private readonly Mock<ILogger> _mockLogger;
    private readonly MyGreeterCsharpServer _generatedServer;

    public UpdateStorageAccountTests()
    {
        _mockLogger = new Mock<ILogger>();
        ServerOptions options = new ServerOptions { EnableAzureSDKCalls = false, SubscriptionId = SubscriptionId };
        _generatedServer = new MyGreeterCsharpServer(options, _mockLogger.Object);
    }

    [Fact]
    public async Task UpdateStorageAccount_Success()
    {
        // Arrange
        ServerCallContext serverCallContext = new TestServerCallContext();
        UpdateStorageAccountRequest request = new UpdateStorageAccountRequest { RgName = ResourceGroupName,
                                                                                SaName = ServiceAccountName,
                                                                                Tags = { { "key", "value" } } };

        // Act
        var response = await _generatedServer.UpdateStorageAccount(request, serverCallContext);

        // Assert
        Assert.NotNull(response);
        Assert.IsType<UpdateStorageAccountResponse>(response);
    }
}

#endif