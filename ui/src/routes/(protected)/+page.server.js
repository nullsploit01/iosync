import { API_BASE_URL } from "@/config";
import { redirect } from "@sveltejs/kit";

export async function load() {
    process.env["NODE_TLS_REJECT_UNAUTHORIZED"] = 0
    const data = await fetch(`${API_BASE_URL}/devices`, {
        method: 'GET',
        credentials: 'include',
    })

    console.log(data);
    

    if (data.status === 401) {
        redirect(303, "/auth/login")
    }

    return await data.json()
}