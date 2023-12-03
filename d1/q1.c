#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>
#include "d.h"
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
    int count = 0;
    while (fgets(buffer, MAX_LENGTH, fileptr)) // Read every line in the text file
    {
        int *nums = malloc(25 * sizeof(int));
        int length = strlen(buffer);
        for (int i = 0; i < length; i++) // now shift through every char in the line
        {
            if (atoi(&(buffer[i])) > 0) // if its a number, add to list
            {
                char c = buffer[i];
                int toAdd = c - '0';
                nums[numOfNums++] = toAdd;
            }
        }
        // printf("%d %d", nums[0], nums[numOfNums - 1]);
        sum = sum + ((nums[0] * 10) + nums[numOfNums - 1]);
        free(nums);
        numOfNums = 0;
        count++;
    }
    printf("%d", sum);
    fclose(fileptr);
    return 0;
}
