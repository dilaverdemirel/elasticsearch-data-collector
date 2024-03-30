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

export class QueryPreviewDataDTO {
    DataSourceId: string = "";
    Query: string = "";

    constructor() {
    }
}

export type ExampleData = {
    [id: string]: string;
}

export class FieldMetaData {
    FieldName: string = "";
    DataType: string = "";
    constructor() {
    }
}

export class QueryPreviewResultDTO {
    ExampleData: ExampleData[] = [];
    MetaDataList: FieldMetaData[] = [];

    constructor() {
    }
}