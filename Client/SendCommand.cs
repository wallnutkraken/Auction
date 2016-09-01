using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;

namespace Client
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
