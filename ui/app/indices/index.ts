export default class Index {
    ID: string = "";
    Name: string = "";
    Description: string = "";
    Valid: boolean = true;
    SqlQuery: string = "";
    Scheduled: boolean = false;
    CronExpression: string = "";
    SyncType: string = "RELOAD_ALL";
    DataSourceId: string = "";
    DocumentField: string = "";

    constructor() {
    }
}

export class ScheduleIndex {
    IndexId: string = "";
    CronExpression: string = "";
    DocumentIdField: string = "";
    SyncType: string = "";

    constructor() {
    }
}