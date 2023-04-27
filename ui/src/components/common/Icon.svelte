<script>
  import {
    IconCirclePlus,
    IconCircleMinus,
    IconCopy,
    IconTrash,
    IconCloud,
    IconRotateClockwise2,
    IconStack,
    IconFolder,
    IconLink,
    IconUnlink,
    IconDeviceFloppy,
    IconPencil,
    IconTool,
    IconBox,
    IconChartCircles,
    IconLayersSubtract,
    IconInfoCircle,
    IconFileText,
    IconAlertCircle,
    IconBrandGit,
    IconDatabase,
    IconSettings,
    IconPlayerPlay,
    IconPlayerStop,
    IconUpload,
    IconTerminal2,
    IconArrowLeft,
    IconArrowRight,
    IconKey,
    IconExternalLink
  } from "@tabler/icons-svelte";

  // using https://devicon.dev
  import Postgres from "./devicons/Postgres.svelte";
  import Redis from "./devicons/Redis.svelte";
  import SQLite from "./devicons/SQLite.svelte";
  import MySQL from "./devicons/MySQL.svelte";
  import MongoDB from "./devicons/MongoDB.svelte";
  import Github from "./devicons/Github.svelte";
  import Docker from "./devicons/Docker.svelte";
  import Go from "./devicons/Go.svelte";
  import Javascript from "./devicons/Javascript.svelte";
  import Python from "./devicons/Python.svelte";
  import Ruby from "./devicons/Ruby.svelte";

  export let type;
  export let size = "md";

  const icons = {
    add: IconCirclePlus,
    remove: IconCircleMinus,
    delete: IconTrash,
    "external-link": IconExternalLink,
    docker: Docker,
    go: Go,
    javascript: Javascript,
    python: Python,
    ruby: Ruby,
    cloud: IconCloud,
    copy: IconCopy,
    restart: IconRotateClockwise2,
    build: IconStack,
    rebuild: IconStack,
    folder: IconFolder,
    link: IconLink,
    unlink: IconUnlink,
    save: IconDeviceFloppy,
    plus: IconCirclePlus,
    edit: IconPencil,
    cube: IconBox,
    left: IconArrowLeft,
    right: IconArrowRight,
    key: IconKey,
    circles: IconChartCircles,
    layers: IconLayersSubtract,
    info: IconInfoCircle,
    "file-text": IconFileText,
    warning: IconAlertCircle,
    error: IconAlertCircle,
    git: IconBrandGit,
    terminal: IconTerminal2,
    database: IconDatabase,
    spanner: IconTool,
    settings: IconSettings,
    start: IconPlayerPlay,
    stop: IconPlayerStop,
    upload: IconUpload,
    postgres: Postgres,
    redis: Redis,
    sqlite: SQLite,
    mysql: MySQL,
    mongodb: MongoDB,
    mongo: MongoDB,
    github: Github,
  };
  const sizeClasses = {
    xs: "w-3 h-3",
    sm: "w-5 h-5",
    md: "w-6 h-6",
    lg: "w-8 h-8",
    xl: "w-10 h-10",
  };
  const sizes = {
    xs: "14",
    sm: "20",
    md: "24",
    lg: "32",
    xl: "38",
  };
</script>

<div class={sizeClasses[size]}>
  <svelte:component this={icons[type]} size={sizes[size]} />
</div>
