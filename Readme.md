# nato The friendly cli tool compact in one single binary file

```
echo "a,b,c,d" | nato loop -s "," -p split -c "echo Value: {{.Value}}, Index: {{.Index}}"
```
