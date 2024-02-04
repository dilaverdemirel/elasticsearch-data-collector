export default class Datasource {
    ID: string = "";
    Name: string = "";
    ConnectionString: string = "";
    MaxPoolSize: number = 10;
    MinIdle: number = 2;
    UserName: string = "";
    DbPassword: string = "";
    DriverName: string = "";

    constructor() {
    }
}
