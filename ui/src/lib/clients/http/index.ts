import { API_BASE_URL } from "@/config/environment";
import axios from "axios";

export const httpClient = axios.create({
    baseURL: API_BASE_URL
}) 