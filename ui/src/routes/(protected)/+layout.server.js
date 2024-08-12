import { redirect } from '@sveltejs/kit';

export function load({ cookies }) {
	console.log(cookies.getAll());
	
	if (!cookies.get('session_id')) {
		throw redirect(303, '/auth/login');
	}
}
