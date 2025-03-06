#if SERVER

using Moq;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using System.Threading.Tasks;
using Serilog;
using MyGreeterCsharp.Api.V1;
using MyGreeterCsharp.Server.Handler;

namespace Server.Tests;

public class ReadStorageAccountTests
{
    private static readonly string SubscriptionId = "test-subscription";
    private static readonly string ResourceGroupName = "test-rg";
    private static readonly string ServiceAccountName = "test-service-account";

    private readonly Mock<ILogger> _mockLogger;
    private readonly MyGreeterCsharpServer _generatedServer;

    public ReadStorageAccountTests()
    {
        _mockLogger = new Mock<ILogger>();
        ServerOptions options = new ServerOptions { EnableAzureSDKCalls = false, SubscriptionId = SubscriptionId };
        _generatedServer = new MyGreeterCsharpServer(options, _mockLogger.Object);
    }

    [Fact]
    public async Task ReadStorageAccount_Success()
    {
        // Arrange
        ServerCallContext serverCallContext = new TestServerCallContext();
        ReadStorageAccountRequest request =
            new ReadStorageAccountRequest { RgName = ResourceGroupName, SaName = ServiceAccountName };

        // Act
        var response = await _generatedServer.ReadStorageAccount(request, serverCallContext);

        // Assert
        Assert.NotNull(response);
        Assert.IsType<ReadStorageAccountResponse>(response);
    }
}

#endif