import { httpClient } from "@/clients/http"

export const authService = {
    login: async (username: string, password: string) => {
        return await httpClient.post('/auth/login', {username, password})
    }
}