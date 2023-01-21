Siec
===

An interpreted Esoteric Programming Language 

Project Setup
-
Each Siec project have to have a code.sc and an inst.sc file <br>
code.sc is where your code belongs<br>
In inst.sc you will write Line Numbers in Base20 that shoud be executed
Each Line Number can only have 2 Digits. This means that your code.sc is limited to jj₂₀ (399₁₀) lines

Benefits
-
* Easy to understand syntax
* Control over Execution via ins.sc
* Variable Names with whitespaces

Commands
-
| Command | Usage                     | Format                         |
|---------|---------------------------|--------------------------------|
| ,       | Println                   | ,\<Text/Variable>              |
| .       | Print                     | .\<Text/Value>                 |
| ^       | Set variable              | ^<var_name>#Value              |
| -       | Subtracts from a variable | -<var_name>#<other_var/number> |
| +       | Adds to a variable        | +<var_name>#<other_var/number> |
| *       | Multiplies a variable     | *<var_name>#<other_var/number> |
| /       | Divides a variable        | /<var_name>#<other_var/number> |
