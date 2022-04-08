<script>
    import { onMount } from "svelte";
    import { loggedIn } from "../store";
    import { navigate } from "svelte-navigator";
    import Result from "./Result.svelte";


    let result = null;

    onMount(async() => {
        const res = await fetch(API_URL + 'api/signout',{
            method: "POST",
            credentials: "include"
        });

        const json = await res.json()
        result = JSON.parse(JSON.stringify(json))

        if (result.success){
            loggedIn.set(false)
            navigate("/")
        }
    });
</script>

<Result result={result} />