interface Api<T>{
    code : number;
    data : T;
    msg: string;
}

interface ApiList<T>{
    code : number;
    Rows : T[];
    Total: number;
}