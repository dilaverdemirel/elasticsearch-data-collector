import axios, { AxiosResponse } from "axios"
import Datasource from "./datasource"

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