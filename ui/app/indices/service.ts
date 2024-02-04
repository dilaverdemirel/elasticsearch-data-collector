import axios, { AxiosResponse } from "axios"
import Index, { ScheduleIndex } from "."

export const getIndexById = (ID: string): Promise<AxiosResponse<Index>> => {
    return axios.get<Index>("http://localhost:8080/indices/" + ID)
}

export const createIndex = (index: Index) => {
    return axios.post("http://localhost:8080/indices",
        index,
        {
            headers: {
                'Content-Type': 'application/json'
            }
        })
}

export const updateIndex = (index: Index) => {
    return axios.put("http://localhost:8080/indices/" + index.ID, index)
}

export const scheduleIndexDataSync = (scheduleIndex: ScheduleIndex) => {
    return axios.put("http://localhost:8080/indices/" + scheduleIndex.IndexId + "/schedule-data-sync", scheduleIndex)
}

export const unscheduleIndexDataSync = (indexId: string) => {
    return axios.delete("http://localhost:8080/indices/" + indexId + "/unschedule-data-sync")
}

export const deleteIndex = (indexId: string) => {
    return axios.delete("http://localhost:8080/indices/" + indexId)
}