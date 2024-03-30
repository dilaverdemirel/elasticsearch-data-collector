import axios, { AxiosResponse } from "axios"
import Datasource, { QueryPreviewDataDTO, QueryPreviewResultDTO } from "./datasource"

export const getDatasources = () => {
    return axios.get("http://localhost:8080/datasources")
}

export const getDatasourceById = (ID: string): Promise<AxiosResponse<Datasource>> => {
    return axios.get<Datasource>("http://localhost:8080/datasources/" + ID)
}

export const createDatasource = (datasource: Datasource) => {
    return axios.post("http://localhost:8080/datasources", datasource)
}

export const updateDatasource = (datasource: Datasource) => {
    return axios.put("http://localhost:8080/datasources/" + datasource.ID, datasource)
}

export const deleteDatasource = (datasourceId: string) => {
    return axios.delete("http://localhost:8080/datasources/" + datasourceId)
}

export const getQueryPreviewData = (previewMetadata: QueryPreviewDataDTO): Promise<AxiosResponse<QueryPreviewResultDTO>> => {
    return axios.post<QueryPreviewResultDTO>("http://localhost:8080/query-meta-data/preview", previewMetadata)
}