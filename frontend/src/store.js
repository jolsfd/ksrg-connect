import { writable } from "svelte/store";

const stored = localStorage.getItem("loggedIn");

export const loggedIn = writable(stored == "true");

loggedIn.subscribe((value) => localStorage.setItem("loggedIn", value))