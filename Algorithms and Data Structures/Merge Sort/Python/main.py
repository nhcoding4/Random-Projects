from typing import Any
import random

# ----------------------------------------------------------------------------------------------------------------------


def main() -> None:
    unsorted_data = [random.randint(0, 101) for _ in range(100)]
    sorted_data: list[int] = merge_sort(unsorted_data)
    print(unsorted_data)
    print(sorted_data)


# ----------------------------------------------------------------------------------------------------------------------


def merge_sort(data: list[Any]) -> list[Any]:
    """Split data down into single parts."""

    if len(data) == 1:
        return data

    left: list[Any] = merge_sort(data[int(len(data) / 2) :])
    right: list[Any] = merge_sort(data[: int(len(data) / 2)])

    return sort_data(left, right)


# ----------------------------------------------------------------------------------------------------------------------


def sort_data(left: list[Any], right: list[Any]) -> list[Any]:
    """Order data from lowest to highest."""

    i: int = 0
    j: int = 0
    sorted_data: list[Any] = []

    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            sorted_data.append(left[i])
            i += 1
        else:
            sorted_data.append(right[j])
            j += 1

    while i < len(left):
        sorted_data.append(left[i])
        i += 1

    while j < len(right):
        sorted_data.append(right[j])
        j += 1

    return sorted_data


# ----------------------------------------------------------------------------------------------------------------------


if __name__ == "__main__":
    main()
