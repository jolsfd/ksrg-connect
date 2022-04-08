<script>
  import Users from "./Users.svelte";
  import Result from "./Result.svelte";

  let deleteUsername = "";
  let result = null;

  async function deleteUser(e) {
    e.preventDefault();

    const res = await fetch(API_URL + "api/admin/delete", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
          deleteUsername,
      }),
    });

    const json = await res.json();
    result = JSON.parse(JSON.stringify(json));
    if (result.success) {
        document.location.reload(true)
    }
  }
  
</script>

<Result result={result} />

<form>
    <div class="mb-3">
      <label for="inputDeleteUsername" class="form-label">Username</label>
      <input
        type="text"
        class="form-control me-2"
        id="inputDeleteUsername"
        aria-describedby="passwordHelp"
        bind:value={deleteUsername}
      />
      <button type="submit" class="btn btn-danger" on:click={deleteUser}
      >Delete</button
    >
    </div>
</form>

<Users showUsername={true}/>