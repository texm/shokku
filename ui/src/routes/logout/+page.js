import { doLogout } from "$lib/api";
import { invalidate } from "$app/navigation";
import { redirect } from "@sveltejs/kit";

export const ssr = false;

export async function load({url, depends}) {
  await doLogout();
  await invalidate("app:load");

  throw redirect(307, "/");
}