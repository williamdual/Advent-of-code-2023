#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>
#include "d.h"
char numberWords[9][6] = {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
int numbers[9] = {1, 2, 3, 4, 5, 6, 7, 8, 9};

int wordToNumber(char *);
int indexOfLatestOccourence(char *, char *, int);

int main()
{
    char *filename = "q1.txt";
    FILE *fileptr = fopen(filename, "r");

    if (fileptr == NULL)
    {
        printf("getoutahere");
        return 1;
    }

    int sum = 0;
    char buffer[MAX_LENGTH];
    int numOfNums = 0;
    int count = 1;
    while (fgets(buffer, MAX_LENGTH, fileptr)) // Read every line in the text file
    {
        int en = -1;
        int ln = -1;
        int eWord = 0;
        int lWord = 0;
        int length = strlen(buffer);
        // search for pure digit numbers
        for (int i = 0; i < length; i++)
        {
            int firstnum = 0;
            if (atoi(&(buffer[i])) > 0) // if its a number, add to list
            {
                if (en == -1)
                {
                    en = i;
                }
                else if (i <= en)
                {
                    en = i;
                }
                if (i >= ln)
                {
                    ln = i;
                }
            }
        }
        // search for word numbers
        for (int i = 0; i < 9; i++)
        {
            char *found = strstr(buffer, numberWords[i]);
            int index = found - buffer;
            if (index < 0)
            {
                continue;
            }
            if (index <= en)
            {
                eWord = wordToNumber(numberWords[i]);
                if (ln == -1)
                {
                    ln = en;
                }
                en = index;
            }
            // check if theres another occourence after it
            int latestOccourence = indexOfLatestOccourence(buffer, numberWords[i], 0);
            if (latestOccourence != index)
            {
                index = latestOccourence;
            }
            if (index >= ln)
            {
                lWord = wordToNumber(numberWords[i]);
                ln = index;
            }
        }
        // now retrieve our 2 numbers
        if (eWord == 0)
        {
            en = buffer[en] - '0';
        }
        else
        {
            en = eWord;
        }
        if (lWord == 0)
        {
            ln = buffer[ln] - '0';
        }
        else
        {
            ln = lWord;
        }
        sum += (en * 10) + ln;

        printf("%d %d %d\n", count, en, ln);

        count++;
    }
    printf("%d", sum);
    fclose(fileptr);
    return 0;
}

int wordToNumber(char *word)
{
    for (int i = 0; i < 9; i++)
    {
        if (strcmp(word, numberWords[i]) == 0)
        {
            return numbers[i];
        }
    }
}

int indexOfLatestOccourence(char *substring, char *word, int carryOver)
{
    char *found = strstr(substring, word);
    int index = found - substring;

    if (index < 0)
    {
        return -1;
    }
    char subsub[MAX_LENGTH];
    strncpy(subsub, &substring[index + 1], strlen(substring));
    index += carryOver + 1; //,might need to +1 this
    int other = indexOfLatestOccourence(subsub, word, index);
    if (other != -1)
    {
        return other;
    }
    else
    {
        return index;
    }
}