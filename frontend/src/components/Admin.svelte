<script>
  import { onMount } from "svelte";
  import { navigate } from "svelte-navigator";

  let data = [];
  let users = [];

  let deleteUsername = "";

  let searchString = "";
  let result = null;

  onMount(async () => {
    const res = await fetch(API_URL + "api/users", {
      method: "GET",
      credentials: "include",
    });

    const json = await res.json();
    if (res.status == 403) {
      navigate("/signin");
    }
    data = JSON.parse(JSON.stringify(json));

    users = data;
  });

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

  function filterUsers() {
    users = data.filter(function (el) {
      return (
        el.username.toLowerCase().indexOf(searchString.toLocaleLowerCase()) > -1 ||
        el.firstName.toLowerCase().indexOf(searchString.toLowerCase()) > -1 ||
        el.lastName.toLowerCase().indexOf(searchString.toLowerCase()) > -1 ||
        el.schoolClass.toLowerCase().indexOf(searchString.toLowerCase()) > -1 ||
        el.age.toString().toLowerCase().indexOf(searchString.toLowerCase()) > -1
      );
    });
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

<div class="d-flex">
  <input
    class="form-control me-2"
    type="search"
    placeholder="Search"
    aria-label="Search"
    bind:value={searchString}
  />
  <button class="btn btn-outline-success" type="submit" on:click={filterUsers}>
    Search
  </button>
</div>

{#each users as item}
    <div class="card mt-2">
        <div class="card-body">
        <h4 class="card-title">{item.username}</h4>
        <h6 class="card-subtitle mb-2 text-muted">{item.firstName}, {item.lastName}</h6>
        <h6 class="card-subtitle mb-2 text-muted">{item.age} years old; class: {item.schoolClass}</h6>
        <p class="card-text">{item.description}</p>
        <p class="card-text">{item.contact}</p>
        </div>
    </div>
{/each}
