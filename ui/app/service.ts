import axios, { AxiosResponse } from "axios"
import StatusStatistic, { RecordStatistic } from "./stats"

export const getStatusStatistic = (): Promise<AxiosResponse<StatusStatistic[]>> => {
    return axios.get<StatusStatistic[]>("http://localhost:8080/indices/sync-daily-status-stats")
}
export const getRecordStatistic = (): Promise<AxiosResponse<RecordStatistic[]>> => {
    return axios.get<RecordStatistic[]>("http://localhost:8080/indices/sync-daily-record-stats")
}
