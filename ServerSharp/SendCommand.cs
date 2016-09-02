namespace ServerSharp
{
    public class SendCommand
    {
        public string Command { get; set; }
        public string[] Args { get; set; }

        public SendCommand() : this("", new string[0])
        {

        }

        public SendCommand(string command, string[] args)
        {
            Command = command;
            Args = args;
        }
    }
}