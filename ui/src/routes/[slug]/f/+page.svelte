<script lang="ts">
  import { onMount } from "svelte";

  type PathEntry = {
    path: string;
    is_directory: boolean;
  };

  let { params } = $props();
  let message = $state("");
  let activePaths = $state([]);
  let entries: PathEntry[] = $state([]);

  const browseFolder = async (paths: string[]) => {
    try {
      const query = new URLSearchParams({
        d: paths.join("/"),
      });

      const res = await fetch(`/api/f?${query.toString()}`);
      const data = await res.json();
      entries = data;
      if (res.ok) {
        message = "call return ok";
        entries = data.entries;
        activePaths = data.paths;
      } else {
        message = `call return non-ok. err=${data}`;
      }
    } catch (error) {
      message = JSON.stringify(error);
    }
  };

  const openDirectory = async (path: string) => {
    const paths = [...activePaths, path];
    await browseFolder(paths);
  };

  const goBack = async (idx: number) => {
    if (idx === -1) {
      await browseFolder([]);
    } else {
      const paths = activePaths.slice(0, idx);
      await browseFolder(paths);
    }
  };

  onMount(async () => {
    await browseFolder([]);
  });
</script>

<div>
  token = {params.slug}
</div>
<div>
  message = {message}
</div>

<div>
  <h1>directory</h1>
  <div style="margin-bottom: 10px;">
    <!-- svelte-ignore a11y_invalid_attribute -->
    <a href="#" onclick={() => goBack(-1)}>root</a>
    {activePaths.length > 0 ? " > " : ""}
    {#each activePaths as activePath, i}
      {#if i + 1 == activePaths.length}
        {`${activePath}`}
      {:else}
        <!-- svelte-ignore a11y_invalid_attribute -->
        <a href="#" onclick={() => goBack(i + 1)}>{activePath}</a>{" > "}
      {/if}
    {/each}
  </div>
  <div>
    {#each entries as entry}
      <div>
        {#if entry.is_directory}
          <button onclick={() => openDirectory(entry.path)}>open</button> - {entry.path}
        {:else}
          {entry.path}
        {/if}
      </div>
    {/each}
  </div>
</div>
