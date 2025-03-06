#if SERVER

using Xunit;
using Moq;
using Serilog;
using Azure.Identity;
using Azure.Core;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using System;
using MyGreeterCsharp.Server.Handler;

namespace Server.Tests;

public class ApiTests
{
    [Fact]
    public void CreatesNewInstance_WhenAzureSDKCallsDisabled()
    {
        // Arrange
        var options = new ServerOptions { 
            EnableAzureSDKCalls = false, 
            SubscriptionId = "test-subscription-id" 
        };

        var mockLogger = new Mock<Serilog.ILogger>();

        // Act
        var server = new MyGreeterCsharpServer(options, mockLogger.Object);

        // Assert
        Assert.NotNull(server);
    }
}

#endif