
# Remarks

I've chosen to complete this assignment using Golang. Golang doesn't have a
maintained xUnit testing framework listed
[here](https://en.wikipedia.org/wiki/List_of_unit_testing_frameworks).
However, the
[stretchr/Testify](https://github.com/stretchr/testify) library
meets most requirements for an xUnit framework.

--Eric

# Introduction 
The simple project designed to let you practice using TDD (Test Driven Development).

# Getting Started
The purpose of the project is to implement a simple conversion function that takes an integer number and converts to a worded string (in whatever output language that you care to use for implementation though the initial test is based on English). The project source is explicitly stubbed out to a minimum implementation that only defines main program, the conversion function declaration, and two unit tests.

# Functional Specification

>## IntegerToWordedString
>
>### Description
>Converts an integer number to its equivalent worded name as string
>such as the integer number `1` being represented as the string `one`.
>
>### Synopsis
>       string IntegerToWordedString(int number);
>
>### Parameters
>   `number`  -   non-negative integer value to be converted
>
>   (return)  -   worded name equivalent of the input `number`
>
>### Notes
>The implementation as provided in the same does not define any requirements
>about how language, culture, or related concerns should / could be
>addressed by the code.

# Developing using the Project

## Project Structure

### `sample/Program.cs`
A simple console app interface for the conversion code. The console app takes a single integer parameter and outputs the converted worded string value on the console.

### `sample/lib/Converter`
The utility class that includes the `IntegerToWordedString` as a public static function.

### `sample-tests/lib/Converter_IntegerToWordedString_Tests`
The complete set of unit tests for the `IntegerToWordedString` function that is the focus of this project. Note the starting set of unit tests contain one test `NotImplementedYet` that may be obsoleted if you implement the complete functionality or the conversionspecification.

## Building and Testing
While the system can be built and tested using different IDEs, this project was created and tested using VSCode. Testing was performed by running the full set of tests in the VSCode terminal window using `dotnet` command -- specifically, executing `dotnet test` after changing the current directory to `sample-tests` folder.

The first execution should result in final "red" state test status like:

><span style='color:red;font-family:monospace;font-weight: bold'>Failed!  - Failed:     1, Passed:     1, Skipped:     0, Total:     2</span>

# Using TDD to Implement the Feature
Before starting to use TDD to implement the feature, we believe that you should take a strong <a href="https://martinfowler.com/bliki/Yagni.html">YAGNI</a> (You Ain't Going to Need It) and emergent design approach to this work. What does this mean? We suggest that you treat each new test case as defining a new specification for the code and that only the current set of test cases is the complete specification. Therefore, you should not design further than the current test cases even if you know more is coming (design upfront thinking). We believe that this will lead to closest experience in letting TDD drive design and refactoring rather than upfront thinking.

This project is set up to allow you to iterate on the design and implementation using red-green-yellow TDD process. The current code of the implementation and test only has stub functionality and testing defined such that code on runs in a Not Implemented exception mode. The first unit test that should test implemented functionality is implemented without its implementation ("red" state).

Create a new git branch to work on your implementation.

To begin doing TDD using the red-green-yellow process, implement the code for IntegerToWordedString that will satisfy the first feature based unit test (the test named `One`). Write and test code until all existing unit tests succeed. Commit the changes to git.

After you are in the "green" state (all unit test pass), add **one** new unit test similar to the existing unit test `One` that will test a different input and expected result. Make sure that all tests pass **except** this one new test -- putting in the "red" state. Commit the changes to git. Write code to implement the functionality verified in the unit test until the test passes again ("green" state), then commit the changes to git.

Review the code and determine whether the code design could (or should) be improved-- if it can be improved ("yellow" state). You may want to consider things like DRY and SOLID principles, as well as run code metrics to help you evaluate the quality of the current design. If so, refactor the design and code making sure that none of changes break the existing tests. Add new tests as part of the refactoring if the new design has edge conditions or other situations that should be verified to maintain the quality of the system. Make sure to current changes to git, as your create new tests and working code.

Repeat this red-green-yellow process until you have fully implemented the full specification for the method.

We recommend that do a Pull Request and merge each iteration branch back into the main code if you want to honor a Continuous Integration (CI) approach to your evolving working design. 
