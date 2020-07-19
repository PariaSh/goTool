# goTools
A tiny tool to 

## Normalize
Normalize ARM instructions

## Remove duplicates
Remove duplicate word from a vocabulary file.

## Remove empty
Remove rows that contain empty fields from a CSV file.

### Output
Two files to be expected.

If the input file is test.csv The output files would be in the same folder with names 
1. test.csv-non-empty, and 
2. test.csv-empty

### How to run it
`./goTool -removeEmpty=true -input=test.csv -indices="1,3,9" hasHeader=true`
- Input file is test.csv
- Indexes of fields to be checked are 1, 3 and 9.

`./goTool -removeEmpty=true -input=test.csv -indices="0"`
- Input file is test.csv
- Index of a field to be checked is 0.

`./goTool -removeEmpty=true -input=test.csv`
- Input file is test.csv
- All fields are to be checked.

## Remove rows with more tokens
Remove rows that contain number of tokens out of limitation.

### Output
Two files to be expected.

If the input file is test.csv The output files would be in the same folder with names 
1. test.csv-non-qualified, and 
2. test.csv-qualified

### How to run it
`./goTool -filter=true -input=test.csv -indices="1,3,9" -hasHeader=true -min=1 -max=512`
- Input file is test.csv
- Indexes of fields to be checked are 1, 3 and 9.
- Minimum length is 1
- Maximum length is 512

`./goTool -filter=true -input=test.csv -indices="1,3,9" -hasHeader=true -min=1`
- Input file is test.csv
- Indexes of fields to be checked are 1, 3 and 9.
- Minimum length is 1

`./goTool -filter=true -input=test.csv -indices="1,3,9" -hasHeader=true`
- Input file is test.csv
- Indexes of fields to be checked are 1, 3 and 9.
- Minimum length is 1


## Merge CSV

### Merge VEX
Merge X86 and ARM CSV files and create fn2fn CSV.

### Merge BinShape and VEX

### How to run
`./goTool -merge=true -f1=f1.csv -f2=f2.csv`

## Shrink file
Randomly select a percentage of lines of a given file to generate a new file

#### How to run

`./goTool -shrink=true -input=fn2fn.csv -p=0.6 -hasHeader -s`

- `input`: input file is "fn2fn.csv"
- `p`: shrink percentage is 0.6
- `h`: keep headers. If true, first line would be appeared in the generated file(s); otherwise, all lines are randomized into the generated file(s)
- `s`: keep both files. If true, keep both fn2fn.csv.0.6 and fn2fn.csv.0.4 

## Rotate folder
Rotate previous dataset folder structure to new dataset folder structure

### Previous structure

{ARCHITECTURE}-{COMPILER}-{LIBRARY}/{VERSION}/-{OPTIMIZATION}/

### New structure

{LIBRARY}-{VERSION}/{ARCHITECTURE}-{COMPILER}-{OPTIMIZATION}{OBFUSCATION}/

## logMatrix
metrics of output from ml

### Input

log file of the format:

```
Epoch 1/33
798/798 [==============================] - 6341s 8s/step - loss: 2.7144 - accuracy: 0.4629 - val_loss: 2.0714 - val_accuracy: 0.4864
Epoch 2/33
798/798 [==============================] - 6511s 8s/step - loss: 2.1010 - accuracy: 0.5019 - val_loss: 1.8767 - val_accuracy: 0.4864
```
### Output

A CSV file with the columns:

```
epoch,total time,time per step,loss,accuracy,val_lost,val_accuracy
```

File name will be the name of the input file appends with `.csv`

### How to run it

```shell script
./goTool -logMatrix=true -input=FILE_NAME
```

- `-input=` accepts a log file name.

e.g.:

```shell script
/goTool -logMatrix=true -input=seq2seq.txt
```

- Input file is `seq2seq.txt`
- Output file is `seq2seq.txt.csv`