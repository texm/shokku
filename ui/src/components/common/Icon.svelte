<script>
  import CirclePlus from "tabler-icons-svelte/icons/CirclePlus.svelte";
  import CircleMinus from "tabler-icons-svelte/icons/CircleMinus.svelte";
  import Copy from "tabler-icons-svelte/icons/Copy.svelte";
  import Trash from "tabler-icons-svelte/icons/Trash.svelte";
  import Cloud from "tabler-icons-svelte/icons/Cloud.svelte";
  import RotateClockwise2 from "tabler-icons-svelte/icons/RotateClockwise2.svelte";
  import Stack from "tabler-icons-svelte/icons/Stack.svelte";
  import Folder from "tabler-icons-svelte/icons/Folder.svelte";
  import Link from "tabler-icons-svelte/icons/Link.svelte";
  import Unlink from "tabler-icons-svelte/icons/Unlink.svelte";
  import DeviceFloppy from "tabler-icons-svelte/icons/DeviceFloppy.svelte";
  import Pencil from "tabler-icons-svelte/icons/Pencil.svelte";
  import Tool from "tabler-icons-svelte/icons/Tool.svelte";
  import Box from "tabler-icons-svelte/icons/Box.svelte";
  import ChartCircles from "tabler-icons-svelte/icons/ChartCircles.svelte";
  import LayersSubtract from "tabler-icons-svelte/icons/LayersSubtract.svelte";
  import InfoCircle from "tabler-icons-svelte/icons/InfoCircle.svelte";
  import FileText from "tabler-icons-svelte/icons/FileText.svelte";
  import AlertCircle from "tabler-icons-svelte/icons/AlertCircle.svelte";
  import BrandGit from "tabler-icons-svelte/icons/BrandGit.svelte";
  import Database from "tabler-icons-svelte/icons/Database.svelte";
  import Settings from "tabler-icons-svelte/icons/Settings.svelte";
  import PlayerPlay from "tabler-icons-svelte/icons/PlayerPlay.svelte";
  import PlayerStop from "tabler-icons-svelte/icons/PlayerStop.svelte";
  import Upload from "tabler-icons-svelte/icons/Upload.svelte";
  import Terminal2 from "tabler-icons-svelte/icons/Terminal2.svelte";
  import ArrowLeft from "tabler-icons-svelte/icons/ArrowLeft.svelte";
  import ArrowRight from "tabler-icons-svelte/icons/ArrowRight.svelte";
  import Key from "tabler-icons-svelte/icons/Key.svelte";
  import ExternalLink from "tabler-icons-svelte/icons/ExternalLink.svelte";

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
    add: CirclePlus,
    remove: CircleMinus,
    delete: Trash,
    "external-link": ExternalLink,
    docker: Docker,
    go: Go,
    javascript: Javascript,
    python: Python,
    ruby: Ruby,
    cloud: Cloud,
    copy: Copy,
    restart: RotateClockwise2,
    build: Stack,
    rebuild: Stack,
    folder: Folder,
    link: Link,
    unlink: Unlink,
    save: DeviceFloppy,
    plus: CirclePlus,
    edit: Pencil,
    cube: Box,
    left: ArrowLeft,
    right: ArrowRight,
    key: Key,
    circles: ChartCircles,
    layers: LayersSubtract,
    info: InfoCircle,
    "file-text": FileText,
    warning: AlertCircle,
    error: AlertCircle,
    git: BrandGit,
    terminal: Terminal2,
    database: Database,
    spanner: Tool,
    settings: Settings,
    start: PlayerPlay,
    stop: PlayerStop,
    upload: Upload,
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
