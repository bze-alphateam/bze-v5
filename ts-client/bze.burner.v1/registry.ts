import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgJoinRaffle } from "./types/burner/tx";
import { MsgFundBurner } from "./types/burner/tx";
import { MsgStartRaffle } from "./types/burner/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/bze.burner.v1.MsgJoinRaffle", MsgJoinRaffle],
    ["/bze.burner.v1.MsgFundBurner", MsgFundBurner],
    ["/bze.burner.v1.MsgStartRaffle", MsgStartRaffle],
    
];

export { msgTypes }