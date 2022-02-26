<script>
    import { loggedIn } from "../store";
    import { navigate } from "svelte-navigator";

    let username = '';
    let password = '';

    let result = null;

    async function signIn(){

        const res = await fetch('http://localhost:8000/api/signin',{
          method: 'POST',
          credentials: 'include',
          body: JSON.stringify({
            username,
            password
          })
        })

        const json = await res.json()
        result = JSON.parse(JSON.stringify(json))
        console.log(result)

        if (result.success) {
          loggedIn.set(true)
          navigate("/explore")
        }
    }
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
<div class="mb-3">
    <label for="inputUsername" class="form-label">Username</label>
    <input type="text" class="form-control" id="inputUsername" bind:value={username}>
  </div>

  <div class="mb-3">
    <label for="inputPassword" class="form-label">Password</label>
    <input type="password" class="form-control" id="inputPassword" bind:value={password}>
  </div>

<button type="submit" class="btn btn-primary" on:click={signIn}>Sign in</button>