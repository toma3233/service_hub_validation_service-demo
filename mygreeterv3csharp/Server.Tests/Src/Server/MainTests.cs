#if SERVER

using System.Threading.Tasks;
using Xunit;
using Moq;
using MyGreeterCsharp.Server.Handler;

namespace Server.Tests;

public class ServerMainCommandTests
{
    [Fact]
    public async Task Main_ShouldInvokeStartCommand()
    {
        // Arrange
        var args = new string[] { "start" };
        var timeout = TimeSpan.FromSeconds(5);

        // Act
        var mainTask = ServerMainCommand.Main(args);
        var timeoutTask = Task.Delay(timeout);

        var completedTask = await Task.WhenAny(mainTask, timeoutTask);

        // Assert
        // If the main task times out, it is the expected behavior, we ignore it. Otherwise we assert failure.
        if (completedTask == timeoutTask) {
            Console.WriteLine("Main command timed out");
        } else {
            var result = await mainTask;
            Assert.Fail("Main command should have timed out");
        }
    }
}

#endif