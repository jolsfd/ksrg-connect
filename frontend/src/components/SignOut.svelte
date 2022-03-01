<script>
    import { onMount } from "svelte";
    import { loggedIn } from "../store";
    import { navigate } from "svelte-navigator";


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

{#if result}
  {#if result.success}
    <div class="alert alert-success" role="alert">
      {result.msg}
    </div>
  {:else}
    <div class="alert alert-danger" role="alert">
      {result.msg}
    </div>
  {/if}
{/if}