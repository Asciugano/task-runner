# TASK-RUNNER

Un programma cli per eseguire script automaticamente(simile a npm)

## Requisiti:
- task-runner
- file di configurazione

## task.yaml
  1. Creare un file con estensione .yaml
  - il nome del file non deve per forza essere **task.yaml**

  2. Se il nome del file non e task o non e' nella root del progetto:
  ```bash
     task-runner --config-path </path/to/your/task.yaml>
```

## Avviare:
  1. **Senza opzioni**
  ```bash
  task-runner [nome_script]
  ```

  2. **Output file**
  ```bash
  task-runner [nome_script] -o <file>
  ```

  3. **Verbose**
  ```bash
  task-runner [nome_script] -v
  ```
