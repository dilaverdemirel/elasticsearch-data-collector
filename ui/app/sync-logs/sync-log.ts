export default class SyncLog {
    ID: string = "";
    IndexId: string = "";
    DocumentCount: number = 0;
    StartDate: Date = new Date();
    EndDate: Date = new Date();
    ExecutionDuration: number = 0;
    Status: string = "";
    StatusMessage: string = "";
    
    constructor() {
    }
}