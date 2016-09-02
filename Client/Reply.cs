using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;

namespace Client
{
    public class Reply
    {
        public string ReplyType { get; set; }
        public string ValueJson { get; set; }

        public override string ToString()
        {
            return $"{ReplyType},{ValueJson}";
        }

    }
}
