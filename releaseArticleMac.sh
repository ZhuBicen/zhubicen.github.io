cd docs
git pull origin master
cd ..
Hugo
case "$(uname -s)" in
   Darwin)
       echo 'Mac OS X'
       ./bin/HugoUtils
     ;;

   Linux)
     echo 'Linux'
     ;;

   CYGWIN*|MINGW32*|MSYS*)
     echo 'MS Windows'
     ./bin/HugoUtils       
     ;;


   *)
     echo 'Other OS' 
     ;;
esac

cd docs
git add *
git commit -m"New Article" -a
git push origin HEAD:master
