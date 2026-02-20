import java.util.Arrays;

public class MergeSort {

    /**
     * Sorts the given integer array in ascending order using merge sort.
     *
     * @param data the array to sort
     */
    public static void sort(int[] data) {
        if (data == null || data.length < 2) {
            return;
        }
        splitAndMerge(data, 0, data.length - 1);
    }

    private static void splitAndMerge(int[] data, int lo, int hi) {
        if (lo >= hi) {
            return;
        }
        int mid = lo + (hi - lo) / 2;
        splitAndMerge(data, lo, mid);
        splitAndMerge(data, mid + 1, hi);
        mergeParts(data, lo, mid, hi);
    }

    private static void mergeParts(int[] data, int lo, int mid, int hi) {
        int[] tmp = Arrays.copyOfRange(data, lo, hi + 1);
        int leftIndex = 0;
        int rightIndex = mid - lo + 1;
        int outputIndex = lo;

        while (leftIndex <= mid - lo && rightIndex <= hi - lo) {
            if (tmp[leftIndex] <= tmp[rightIndex]) {
                data[outputIndex++] = tmp[leftIndex++];
            } else {
                data[outputIndex++] = tmp[rightIndex++];
            }
        }
        while (leftIndex <= mid - lo) {
            data[outputIndex++] = tmp[leftIndex++];
        }
        while (rightIndex <= hi - lo) {
            data[outputIndex++] = tmp[rightIndex++];
        }
    }

    public static void main(String[] args) {
        int[] numbers = {64, 34, 25, 12, 22, 11, 90};

        System.out.println("Before: " + Arrays.toString(numbers));
        sort(numbers);
        System.out.println("After:  " + Arrays.toString(numbers));
    }
}
