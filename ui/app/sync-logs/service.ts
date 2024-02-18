import axios, { AxiosResponse } from "axios"
import SyncLog from "./sync-log"

export const findByIndexId = (IndexId: string): Promise<AxiosResponse<SyncLog>> => {
    return axios.get<SyncLog>("http://localhost:8080/sync-logs?IndexId=" + IndexId + "&size=10")
}
