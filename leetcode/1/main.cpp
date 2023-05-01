#include <algorithm>
#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
public:
  vector<int> twoSum(vector<int> &nums, int target) {
    // create a map to store the value and index
    unordered_map<int, int> map;

    // create array to store the result
    vector<int> result;

    for (int i = 0; i < nums.size(); i++) {
      int element = nums[i];
      auto map_res = map.find(target - element);
      if (map_res != map.end()) {
        result.push_back(i);
        result.push_back(map_res->second);
        return result;
      }
      map.insert({element, i});
    }

    return result;
  }
};
