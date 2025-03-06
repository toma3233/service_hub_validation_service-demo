using Xunit;
using Moq;
using Serilog;
using Grpc.Net.Client;
using MyGreeterCsharp.Api.V1;

namespace Api.Tests;
public class ClientFactoryTests
{
    [Fact]
    public void NewClient_ShouldReturnMyGreeterCsharpClient()
    {
        // Arrange
        var remoteAddr = "localhost:5000";
        var mockLogger = new Mock<ILogger>();

        // Act
        var client = ClientFactory.NewClient(remoteAddr, mockLogger.Object);

        // Assert
        Assert.NotNull(client);
        Assert.IsType(typeof(MyGreeterCsharp.Api.V1.MyGreeterCsharp.MyGreeterCsharpClient), client);
    }
}
