import { useState } from "react"
import React from "react"
export enum ApiStatus{
    Loading,
    Success,
    ErrorUnauthorised,
    Error,
    RefreshingToken,
    Retrying,
}

interface IApiData {
    status: ApiStatus,
    error: any,
    data: any,
}

export const useApi = (url: string, body : {}) => {
    const [retryToggle, setRetryToggle] = useState(false)
    const [data, setData] = useState<IApiData>({
        status: ApiStatus.Loading ,
        error: null,
        data: null
    })

}