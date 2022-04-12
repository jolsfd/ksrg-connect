<script>
  import { onMount } from "svelte";
  import { navigate } from "svelte-navigator";
  import Loading from "./Loading.svelte";
  import User from "./User.svelte";

  let data = [];
  let users = [];

  let loading = true;

  let searchString = "";

  export let showUsername = false;

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
    loading = false;
  });

  function filterUsers() {
    users = data.filter(function (el) {
      return (
        el.firstName.toLowerCase().indexOf(searchString.toLowerCase()) > -1 ||
        el.lastName.toLowerCase().indexOf(searchString.toLowerCase()) > -1 ||
        el.schoolClass.toLowerCase().indexOf(searchString.toLowerCase()) > -1 ||
        el.age.toString().toLowerCase().indexOf(searchString.toLowerCase()) >
          -1 ||
        el.username.toLowerCase().indexOf(searchString.toLowerCase()) > -1
      );
    });
  }

  function handleKeydown(e) {
    if (e.keyCode === 13) {
      filterUsers();
    }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

<div class="d-flex">
  <input
    class="form-control me-2"
    type="search"
    placeholder="Search"
    aria-label="Search"
    bind:value={searchString}
  />
  <button class="btn btn-outline-success" type="submit" on:click={filterUsers}>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="16"
      height="16"
      fill="currentColor"
      class="bi bi-search-heart"
      viewBox="0 0 16 16"
    >
      <path
        d="M6.5 4.482c1.664-1.673 5.825 1.254 0 5.018-5.825-3.764-1.664-6.69 0-5.018Z"
      />
      <path
        d="M13 6.5a6.471 6.471 0 0 1-1.258 3.844c.04.03.078.062.115.098l3.85 3.85a1 1 0 0 1-1.414 1.415l-3.85-3.85a1.007 1.007 0 0 1-.1-.115h.002A6.5 6.5 0 1 1 13 6.5ZM6.5 12a5.5 5.5 0 1 0 0-11 5.5 5.5 0 0 0 0 11Z"
      />
    </svg>
    Search
  </button>
</div>

<Loading {loading} />

{#each users as item}
  <User
    firstName={item.firstName}
    lastName={item.lastName}
    schoolClass={item.schoolClass}
    age={item.age}
    description={item.description}
    contact={item.contact}
    username={item.username}
    {showUsername}
  />
{/each}
